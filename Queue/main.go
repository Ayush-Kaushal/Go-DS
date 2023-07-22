package main

import "fmt"

type Queue struct {
	items []int
	size  int
}

// add: adds a value at the end
func (q *Queue) add(value int) {
	q.items = append(q.items, value)
	q.size++
}


// remove: removes the value at the front and returns the removed value
func (q *Queue) remove() int {
	if q.size == 0 {
		panic("Queue is empty!")
	}

	removedItem := q.items[0]
	q.items = q.items[1:]
	q.size--

	return removedItem
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.add(100)
	myQueue.add(200)
	myQueue.add(300)
	myQueue.add(400)
	fmt.Println(myQueue.items)

	fmt.Println(myQueue.remove())
	fmt.Println(myQueue.items)

	mtQ := Queue{}
	fmt.Println(mtQ.remove())
}