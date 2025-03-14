package input

import "fmt"

func GetInfo() {

	println("Enter Your Name: ")
	var Name int
	fmt.Scanln(&Name)
	println("Enter Your Age: ")
	var Age int
	fmt.Scanln(&Age)
}
