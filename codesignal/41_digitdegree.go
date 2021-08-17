

func digitDegree(n int) (times int) {
	value := strconv.Itoa(n)
	if len(value) < 2 {
		return times
	}
	for {
		sum := 0
		for _, v := range value {
			val, _ := strconv.ParseInt(string(v), 10, 32)
			sum += int(val)
		}
		times++
		value = strconv.Itoa(sum)
		fmt.Println(times, sum, value)
		if len(value) < 2 {
			break
		}

	}

	return times
}

	
