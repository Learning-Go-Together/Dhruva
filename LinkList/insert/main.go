package insert

import (
	"fmt"
)

type Node struct {
	data int64
	next *Node
}
type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertAtLastPosition(data int64) {
	if ll.head == nil {
		ll.head = &Node{data: data, next: nil}
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{data: data}
}
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d --> ", current.data)
		current = current.next
	}
	println("Null")

}
func (ll *LinkedList) InsertAtFirstPosition(data int64) {
	// if ll.head == nil {
	// 	ll.head = &Node{data: data, next: nil}
	// 	return
	// }
	ll.head = &Node{data: data, next: ll.head}
}
func (ll *LinkedList) Count() int {

	current := ll.head
	count := 0
	for current != nil {
		current = current.next
		count += 1
	}
	return count

}
func (ll *LinkedList) DeleteAtSpecificPosition(index int) {
	current := ll.head
	if index < 0 {
		return
	}
	if index >= ll.Count()-1 {
		ll.head = ll.head.next
		return
	}
	for i := 0; i <= index-1; i++ {
		current = current.next
	}
	if current.next == nil {
		current = current.next
	}
	current = current.next.next

}

func (ll *LinkedList) InsertAtSpecificPosition(index int, data int64) {
	current := ll.head
	if index < 0 || index > ll.Count() {
		return
	}
	if index == 0 {
		ll.InsertAtFirstPosition(data)
	}
	for i := 0; i <= index-1; i++ {
		current = current.next
	}
	newNode := &Node{data: data, next: nil}
	temp := current.next
	current.next = newNode
	newNode.next = temp
}
func (ll *LinkedList) ReverseLinkList() {
	curr := ll.head
	var prev *Node
	if curr == nil {
		return
	}
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
}
