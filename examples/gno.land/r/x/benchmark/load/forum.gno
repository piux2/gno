package load

import (
	"strconv"

	"gno.land/p/demo/avl"
)

// 100 boards on the forum
const BOARDS_NUM int = 10

var boards avl.Tree

type Board struct {
	id    int
	posts *avl.Tree
}

func (b *Board) AddPost(title, content string) {
	n := b.posts.Size()
	p := Post{n, title, content}
	b.posts.Set(strconv.Itoa(n), p)
}

func (b *Board) GetPost(id int) (Post, bool) {
	p, ok := b.posts.Get(strconv.Itoa(id))
	if ok {
		return p.(Post), ok
	} else {
		return Post{}, ok
	}
}

type Post struct {
	id      int
	title   string
	content string
}

// initate
func init() {
	for i := 0; i < BOARDS_NUM; i++ {
		newBoard()
	}
}

func newBoard() {
	n := boards.Size()
	posts := avl.NewTree()
	b := &Board{n, posts}
	boards.Set(strconv.Itoa(n), b)
}

// post to all boards.
func AddPost(title, content string) {
	for i := 0; i < boards.Size(); i++ {
		boardId := strconv.Itoa(i)
		b, ok := boards.Get(boardId)
		if ok {
			b.(*Board).AddPost(title, content)
		}
	}
}

func GetPost(boardId, postId int) string {
	b, ok := boards.Get(strconv.Itoa(boardId))
	var res string

	if ok {
		p, ok := b.(*Board).GetPost(postId)
		if ok {
			res = p.title + "," + p.content
		}
	}
	return res
}

func GetPostSize(boardId int) int {
	b, ok := boards.Get(strconv.Itoa(boardId))
	var res int

	if ok {
		res = b.(*Board).posts.Size()
	} else {
		res = -1
	}

	return res
}

func GetBoardSize() int {
	return boards.Size()
}
