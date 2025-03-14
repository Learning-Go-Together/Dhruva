package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	mydict := make(map[string]int)
	mydict["Dhruva"] = 1
	mydict["Shivang"] = 1
	mydict["Muffin"] = 1
	delete(mydict, "Shivang")
	fmt.Println(mydict)
	// elem, ok := mydict["Dhruva"]
	// if ok {
	// 	fmt.Println("Yes its present in the list")
	// }
	// fmt.Println(elem)
	if _, ok := mydict["Akshat"]; !ok {
		fmt.Println("Error: Akshat is in the dictionary")
	}

}
