func differentSymbolsNaive(s string) int {
	numbers := make(map[rune]bool, len(s))
	for _, v := range s {
		numbers[v] = true
	}

	return len(numbers)
}
