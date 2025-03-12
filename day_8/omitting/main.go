package main

func maxMessages(thresh int) int {
	count := 0
	cost := 0
	for i := 0; ; i++ {
		cost += 100 + i
		if cost > thresh {
			break
		}
		count++

	}
	return count
}
