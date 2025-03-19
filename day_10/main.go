package main

import (
	pingpong "github.com/Learning-Go-Together/Dhruva/day_10/pingPong"
)

func main() {
	// fmt.Println("Hello world")
	// mydict := make(map[string]int)
	// mydict["Dhruva"] = 1
	// mydict["Shivang"] = 1
	// mydict["Akshat"] = 1
	// delete(mydict, "Shivang")
	// fmt.Println(mydict)
	// elem, ok := mydict["Dhruva"]
	// if ok {
	// 	fmt.Println("Yes its present in the list")
	// }
	// fmt.Println(elem)
	// if _, ok := mydict["Akshat"]; !ok {
	// 	delete(mydict, "Akshat")
	// 	fmt.Println("Error: Akshat is in the dictionary")
	// 	fmt.Println(mydict)
	// }
	// var p *int
	// p = 10
	// fmt.Println(p)
	// ch := channel.MakeChannel(10)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	pingpong.Test(2)
	pingpong.Test(3)
	pingpong.Test(4)

}
