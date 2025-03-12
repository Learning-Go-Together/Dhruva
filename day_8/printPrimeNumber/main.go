package printprimenumber

import (
	"fmt"
)

func PrintPrimes(max int) {
	for i := 2; i <= max; i++ {
		if max%i == 0 {
			return fmt.Sprintf("%v is not a Prime Number", max)
		}
	}
	fmt.Sprintf("%v is the Prime Number", max)
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	PrintPrimes(max)
	fmt.Println("===============================================================")
}
