package linked

import (
	"fmt"
	"time"
)

const SLEEP = 500

type Field struct {
	Cell string
	Rent int
}

func (f *Field) Action() {
	fmt.Printf("%+v\n", f)
	time.Sleep(SLEEP * time.Millisecond)
}

type Node struct {
	Value    *Field
	Next     *Node
	Previous *Node
}

func (li *Node) Init() *Node {
	return &Node{
		Value:    nil,
		Next:     nil,
		Previous: nil,
	}
}

func (li *Node) Insert(data *Field) (*Node, int) {
	if li.Value == nil {
		li.Value = data
		return li, 0
	}
	li.Next = &Node{
		Value:    data,
		Next:     li.Next,
		Previous: li.Previous,
	}
	return li.Next, 0
}

func (li *Node) Print() {
	//fmt.Printf("%v\n", li)
	fmt.Printf("%+v\n", li)
}

var fuse = 9

func (li *Node) Traverse() {
	if fuse < 1 {
		return
	}
	fuse--
	//if li.Value != nil {
	li.Value.Action()
	if li.Next != nil {
		li.Next.Traverse()
	}
}
