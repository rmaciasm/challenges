func shapeArea(n int) (area int) {
	if n == 1 {
		return 1
	}

	area = 1

	for n-1 != 0 {
		area += (n - 1) * 4
		n--
	}

	return
}