func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if weight1+weight2 <= maxW {
		return value1 + value2
	} else if weight1 <= maxW && (value1 >= value2 || weight2 > maxW) {
		return value1
	} else if weight2 <= maxW && (value1 <= value2 || weight1 > maxW) {
		return value2
	}

	return 0
}
