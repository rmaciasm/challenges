func evenDigitsOnly(n int) bool {
	str := strconv.Itoa(n)
	for i := range str {
		if str[i]%2 != 0 {
			return false
		}
	}

	return true
}