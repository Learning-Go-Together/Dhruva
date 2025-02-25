package age

func VotingRules() [1]string {
	_, res := ValidAge(21)
	rules := [1]string{"Rules: 1. Sanitise your hands \n 2. Wear mask \n 3. Maintain some distance for safety "}
	if res {
		return rules
	}
	return rules
}
