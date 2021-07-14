func palindromeRearranging(inputString string) bool {
	lenght := len(inputString)
	lettersMap := make(map[string]int, 0)
	for _, v := range inputString {
		lettersMap[string(v)]++
	}
	count := 0
	for _, v := range lettersMap {
		if v%2 != 0 {
			count++
		}
	}

	return (count == 0 && lenght%2 == 0) || (count < 2 && lenght%2 == 1)
}
