func alternatingSums(a []int) (res []int) {
	res = make([]int, 2)
	for k, v := range a {
		if k%2 == 1 {
			res[1] += v
		} else {
			res[0] += v
		}
	}

	return
}
