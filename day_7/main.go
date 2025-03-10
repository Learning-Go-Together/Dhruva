package main

import (
	"fmt"
	"math"

	errorinterfaces2 "github.com/Learning-Go-Together/Dhruva/day_7/errorInterfaces2"
)

func main() {
	// practice.ConvertStringToInt("Dhruva")

	value, err := errorinterfaces2.Divide(math.Inf(-1), math.Inf(1))
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(value)
}
