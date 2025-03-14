package main

import "fmt"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case "pro":
		return messages[:], nil
	case "free":
		return messages[:2], nil
	default:
		return nil, fmt.Errorf("unsupported plan")
	}

}
