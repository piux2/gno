package async

import (
	"sync/atomic"

	"github.com/gnolang/gno/tm2/pkg/errors"
)

//----------------------------------------
// Routine

func Routine(fn func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fn()
		close(done)
	}()
	return done
}

//----------------------------------------
// Task

// val: the value returned after task execution.
// err: the error returned during task completion.
// abort: tells Parallel to return, whether or not all tasks have completed.
type Task func(i int) (val any, err error, abort bool)

type TaskResult struct {
	Value any
	Error error
}

type TaskResultCh <-chan TaskResult

type taskResultOK struct {
	TaskResult
	OK bool
}

type TaskResultSet struct {
	chz     []TaskResultCh
	results []taskResultOK
}

func newTaskResultSet(chz []TaskResultCh) *TaskResultSet {
	return &TaskResultSet{
		chz:     chz,
		results: make([]taskResultOK, len(chz)),
	}
}

func (trs *TaskResultSet) Channels() []TaskResultCh {
	return trs.chz
}

func (trs *TaskResultSet) LatestResult(index int) (TaskResult, bool) {
	if len(trs.results) <= index {
		return TaskResult{}, false
	}
	resultOK := trs.results[index]
	return resultOK.TaskResult, resultOK.OK
}

// NOTE: Not concurrency safe.
// Writes results to trs.results without waiting for all tasks to complete.
func (trs *TaskResultSet) Reap() *TaskResultSet {
	for i := range trs.results {
		trch := trs.chz[i]
		select {
		case result, ok := <-trch:
			if ok {
				// Write result.
				trs.results[i] = taskResultOK{
					TaskResult: result,
					OK:         true,
				}
			}
			// else {
			// We already wrote it.
			// }
		default:
			// Do nothing.
		}
	}
	return trs
}

// NOTE: Not concurrency safe.
// Like Reap() but waits until all tasks have returned or panic'd.
func (trs *TaskResultSet) Wait() *TaskResultSet {
	for i := range trs.results {
		trch := trs.chz[i]
		result, ok := <-trch
		if ok {
			// Write result.
			trs.results[i] = taskResultOK{
				TaskResult: result,
				OK:         true,
			}
		}
		// else {
		// We already wrote it.
		// }
	}
	return trs
}

// Returns the firstmost (by task index) error as
// discovered by all previous Reap() calls.
func (trs *TaskResultSet) FirstValue() any {
	for _, result := range trs.results {
		if result.Value != nil {
			return result.Value
		}
	}
	return nil
}

// Returns the firstmost (by task index) error as
// discovered by all previous Reap() calls.
func (trs *TaskResultSet) FirstError() error {
	for _, result := range trs.results {
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

//----------------------------------------
// Parallel

// Run tasks in parallel, with ability to abort early.
// Returns ok=false iff any of the tasks returned abort=true.
// NOTE: Do not implement quit features here.  Instead, provide convenient
// concurrent quit-like primitives, passed implicitly via Task closures. (e.g.
// it's not Parallel's concern how you quit/abort your tasks).
func Parallel(tasks ...Task) (trs *TaskResultSet, ok bool) {
	taskResultChz := make([]TaskResultCh, len(tasks)) // To return.
	taskDoneCh := make(chan bool, len(tasks))         // A "wait group" channel, early abort if any true received.
	numPanics := new(int32)                           // Keep track of panics to set ok=false later.
	ok = true                                         // We will set it to false iff any tasks panic'd or returned abort.

	// Start all tasks in parallel in separate goroutines.
	// When the task is complete, it will appear in the
	// respective taskResultCh (associated by task index).
	for i, task := range tasks {
		taskResultCh := make(chan TaskResult, 1) // Capacity for 1 result.
		taskResultChz[i] = taskResultCh
		go func(i int, task Task, taskResultCh chan TaskResult) {
			// Recovery
			defer func() {
				if pnk := recover(); pnk != nil {
					atomic.AddInt32(numPanics, 1)
					// Send panic to taskResultCh.
					taskResultCh <- TaskResult{nil, errors.Wrap(pnk, "Panic in task")}
					// Closing taskResultCh lets trs.Wait() work.
					close(taskResultCh)
					// Decrement waitgroup.
					taskDoneCh <- false
				}
			}()
			// Run the task.
			val, err, abort := task(i)
			// Send val/err to taskResultCh.
			// NOTE: Below this line, nothing must panic/
			taskResultCh <- TaskResult{val, err}
			// Closing taskResultCh lets trs.Wait() work.
			close(taskResultCh)
			// Decrement waitgroup.
			taskDoneCh <- abort
		}(i, task, taskResultCh)
	}

	// Wait until all tasks are done, or until abort.
	// DONE_LOOP:
	for range tasks {
		abort := <-taskDoneCh
		if abort {
			ok = false
			break
		}
	}

	// Ok is also false if there were any panics.
	// We must do this check here (after DONE_LOOP).
	ok = ok && (atomic.LoadInt32(numPanics) == 0)

	return newTaskResultSet(taskResultChz).Reap(), ok
}
