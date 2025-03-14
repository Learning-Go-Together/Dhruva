package variadic

func sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}
