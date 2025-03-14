package main

import "github.com/Learning-Go-Together/Dhruva/day_9/input"

func printSlice(slice ...string) []string {
	return slice
}

func Reduce(slice []int, reducer func(int, int) int, initial int) int {
	accumulator := initial
	for _, value := range slice {
		accumulator = reducer(accumulator, value)
	}
	return accumulator
}

func main() {
	// slice := []string{"Dhruva", "Shivang", "akshat"}
	// primes := []int{2, 3, 5, 7, 11, 13}
	// var ds []int = make([]int, 5)
	// mySlice := primes[:]
	// fmt.Println(mySlice, ds)

	// fmt.Println(printSlice(slice...))
	// slice = append(slice, "muffin")
	// fmt.Println(printSlice(slice...))

	// numbers := []int{1, 2, 3, 4, 5}

	// Using Reduce function to sum numbers
	// sum := Reduce(numbers, func(acc, num int) int {
	// 	return acc + num
	// }, 0)

	// fmt.Println("Sum:", sum) // Output: Sum: 15

	// rows := make([]int, 5, 9)
	// slice := []int{1, 2, 3}
	// array := [1]int{}
	// fmt.Println(rows)
	input.GetInfo()

}
