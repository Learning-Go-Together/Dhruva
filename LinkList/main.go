package main

import (
	"fmt"
	"linklist/insert"
)

func main() {
	fmt.Println("Welcome To Link List World")
	ll := insert.LinkedList{}
	ll.InsertAtLastPosition(6)
	ll.InsertAtLastPosition(7)
	ll.PrintList()
	ll.InsertAtFirstPosition(8)
	ll.PrintList()
	// ll.DeleteAtSpecificPosition(5)
	ll.InsertAtSpecificPosition(2, 10)
	ll.PrintList()

}
