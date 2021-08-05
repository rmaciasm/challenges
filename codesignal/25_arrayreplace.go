func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for k, v := range inputArray {
		if v == elemToReplace {
			inputArray[k] = substitutionElem
		}
	}

	return inputArray
}