func allLongestStrings(inputArray []string) []string {
	max := -1
	returnMap := make(map[int][]string)
	for _, v := range inputArray {
		returnMap[len(v)] = append(returnMap[len(v)], v)
		if len(v) > max {
			max = len(v)
		}
	}

	return returnMap[max]
}
