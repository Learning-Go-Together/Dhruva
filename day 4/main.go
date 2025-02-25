package main

import (
	processsingorder "day4/processsing_order"
	"fmt"
)

func main() {
	// addstring.Print("Hi", "World!")
	// v := subscriptionfinder.GetMonthlyPrice("muffin")
	// switch v.(type) {
	// case string:
	// 	fmt.Println(v)
	// case int:
	// 	fmt.Println(v)
	// }
	// fmt.Println(billingsystem.CalculateFinalBill(2, 10))
	fmt.Println(processsingorder.PlaceOrder("1", 2, 500))
}
