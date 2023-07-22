package main

import "fmt"

type Stack struct{
	items []int
	size int
}

// Push: adds the value at the end
func (s *Stack) Push(val int){
	s.items = append(s.items, val)
	s.size++;
}

// Pop: removes the value at the end and returns the removed value
func (s *Stack) Pop() int{
	topItem := s.items[s.size - 1]
	s.items = s.items[: s.size - 1]
	s.size--

	return topItem
}

func main() {
	myStack := Stack{}

	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack.items)

	myStack.Pop()
	fmt.Println(myStack.items)
}