func adjacentElementsProduct(inputArray []int) int {
	if len(inputArray) < 2 {
		return -1000
	}

	max := inputArray[0] * inputArray[1]
	prod := adjacentElementsProduct(inputArray[1:])

	if prod > max {
		max = prod
	}

	return max

}