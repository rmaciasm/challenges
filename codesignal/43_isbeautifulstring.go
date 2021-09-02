func isBeautifulString(inputString string) bool {
	m := make(map[rune]int)
	for _, letter := range inputString {
		m[letter]++
	}

	for letter := 'b'; letter <= rune('a'+len(m)); letter++ {
		previous, exists := m[letter-1]

		if !exists || previous < m[letter] {
			return false
		}
	}
	return true
}
   