package fizzbuzz

import (
	"fmt"
)

//	func Fizzbuzz() {
//		for i := 0; i <= 100; i++ {
//			if i%3 == 0 {
//				fmt.Println("fizz")
//				continue
//			}
//			if i%5 == 0 {
//				fmt.Println("buzz")
//				continue
//			}
//			fmt.Println(i)
//		}
//	}
func Fizzbuzz2() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)

		}
	}
}
