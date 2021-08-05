func extractEachKth(inputArray []int, k int) []int {
	newArray := make([]int, 0)
	for p, v := range inputArray {
		if (p+1)%k != 0 {
			newArray = append(newArray, v)
		}
	}

	return newArray
}
