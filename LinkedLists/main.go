package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
	length int
}

func (l *LinkedList) addFirst(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++

}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0{
		fmt.Println("Empty list!")
		return
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value{
		if previousToDelete.next.next == nil{
			fmt.Println("No such value present in the list")
			return
		}

		previousToDelete = previousToDelete.next
	}


	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l LinkedList) printListData(){
	toPrint := l.head

	for l.length != 0{
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Println()
}

func main()  {
	myList := LinkedList{}

	node1 := &Node{data:48}
	node2 := &Node{data:18}
	node3 := &Node{data:16}
	node4 := &Node{data:11}
	node5 := &Node{data:7}
	node6 := &Node{data:2}

	myList.addFirst(node1)
	myList.addFirst(node2)
	myList.addFirst(node3)
	myList.addFirst(node4)
	myList.addFirst(node5)
	myList.addFirst(node6)

	myList.printListData()

	myList.deleteWithValue(16)
	myList.printListData()

	myList.deleteWithValue(10)

	emptyList := LinkedList{}
	emptyList.deleteWithValue(16)
}