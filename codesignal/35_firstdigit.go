func firstDigit(inputString string) (found string) {
	for _, v := range inputString {
		if v >= 48 && v <= 57 {
			return string(v)
		}
	}
	return found
}
	