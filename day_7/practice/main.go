package practice

import (
	"fmt"
	"strconv"
)

// type Error interface {
// 	Error() string
// }

// type Laptop struct {
// 	modelNumber string
// }

func ConvertStringToInt(input string) {
	// fmt.Println("Hello world")
	result, err := convertStringToInt(input)
	if err != nil {
		fmt.Println("Error Occured", result)
		fmt.Printf("err: %v", err)
		return
	}
	// print("succesfull")
	fmt.Println("converted Integer", result)
}

func convertStringToInt(input string) (int, error) {
	c, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Print("conversion Successfull")
	return c, nil
}
