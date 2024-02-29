package main

import (
	"fmt"
	"sync"
)

type Task interface {
	Execute() error
	Key() string
}

type Result struct {
	worker   int
	taskName string
	err      error
}

func monitor(resCh <-chan Result, numTasks int) {
	i := 0
	for v := range resCh {
		i++
		p := (100 * i) / numTasks

		fmt.Printf("Processing: %d%% complete\n", p)
		if v.err != nil {
			fmt.Printf("\nWorker %d finished job %s with error: %s\n", v.worker, v.taskName, v.err)
		}

	}
	fmt.Println()
}

func worker(id int, tasks <-chan Task, taskWg *sync.WaitGroup, resCh chan<- Result) {
	for t := range tasks {
		err := t.Execute()
		res := Result{id, t.Key(), err}
		resCh <- res
		taskWg.Done()
	}
}
