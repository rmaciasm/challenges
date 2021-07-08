func isLucky(n int) bool {
	s := strconv.Itoa(n)
	sum1 := 0
	for i := 0; i < len(s)/2; i++ {
		v, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		sum1 += v

	}
	sum2 := 0
	for j := len(s) / 2; j < len(s); j++ {
		v, err := strconv.Atoi(string(s[j]))
		if err != nil {
			return false
		}
		sum2 += v
	}

	return sum1 == sum2
}
	