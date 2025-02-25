package subscriptionfinder

func GetMonthlyPrice(tier string) interface{} {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return "tier Not Found"
	}
}
