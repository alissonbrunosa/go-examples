package stack

import "fmt"

type node struct {
	Data     interface{}
	Previous *node
}

type Stack struct {
	top    *node
	length int
}

func New() *Stack {
	return &Stack{}
}

func (this *Stack) Push(data interface{}) {
	if this.top == nil {
		this.top = &node{Data: data}
	} else {
		newTop := &node{Data: data}
		newTop.Previous = this.top
		this.top = newTop
	}

	this.length += 1
}

func (this *Stack) Pop() interface{} {
	top := this.top
	this.top = top.Previous
	this.length -= 1
	return top.Data
}

func (this *Stack) Top() interface{} {
	return this.top.Data
}

func (this *Stack) Print() {
	current := this.top
	for current != nil {
		fmt.Printf("Value is %v \n", current.Data)
		current = current.Previous
	}
}

func (this *Stack) IsEmpty() bool {
	return this.length == 0
}

func (this *Stack) Length() int {
	return this.length
}
