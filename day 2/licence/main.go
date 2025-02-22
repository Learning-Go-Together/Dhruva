package age

import (
	"fmt"
)

func ValidAge(age int) (string, bool) {
	if age >= 18 {
		return "You are eligible for voting of 2025", true

	} else {
		return fmt.Sprintf("You are not eligible wait for %d years", 18-age), false
	}
}
