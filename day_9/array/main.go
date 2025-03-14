package array

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	reminder_messages := [3]string{primary, secondary, tertiary}
	cost := [3]int{len(primary), len(secondary) + len(primary), len(tertiary) + len(primary) + len(secondary)}
	return reminder_messages, cost
}
