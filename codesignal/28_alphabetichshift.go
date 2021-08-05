func alphabeticShift(inputString string) (newString string) {
	for _, v := range inputString {
		if v == 122 {
			v = 96
		}
		newString += string(v + 1)
	}
	return
}