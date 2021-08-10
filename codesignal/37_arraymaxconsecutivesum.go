func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := -10000

	for key := 0; key <= len(inputArray)-k; key++ {
		sum := 0
		for i := key; i < key+k; i++ {
			sum += inputArray[i]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
