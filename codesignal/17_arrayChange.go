func arrayChange(inputArray []int) int {
	output := 0
	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i] >= inputArray[i+1] {
			moves := inputArray[i] - inputArray[i+1] + 1
			inputArray[i+1] += moves
			output += moves
		}
	}
	return output
}
