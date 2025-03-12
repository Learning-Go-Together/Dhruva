package loops

func bulkSend(numMessages int) float64 {
	total_cost := 0.0
	for i := 0; i < numMessages; i++ {
		total_cost += 1.00 + 0.01*float64(i)
	}
	return float64(total_cost)
}
