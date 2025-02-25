package billingsystem

func CalculateFinalBill(costPerMessage float64, numMesages int) float64 {
	finalBill := CalculateBaseBill(costPerMessage, numMesages) - CalculateDiscountRate(numMesages)
	return finalBill
	// return costPerMessage * float64(numMesages)
}

func CalculateDiscountRate(messagesSent int) float64 {
	if messagesSent >= 5000 {
		return 0.20
	}
	if messagesSent >= 1000 {
		return 0.10
	}
	return 0
}

func CalculateBaseBill(costPerMessage float64, messageSent int) float64 {
	return costPerMessage * float64(messageSent)
}
