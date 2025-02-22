package age

func VotingRules() [1]string {
	_, res := ValidAge(21)

	if res {
		return [1]string{"Rules: 1. Sanitise your hands \n 2. Wear mask \n 3. Maintain some distance for safety "}
	}
	return [1]string{"You are not eligible to vote, so these rules do not apply to you."}
}
