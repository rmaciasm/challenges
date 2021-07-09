func sortByHeight(a []int) []int {
	for p := 0; p < len(a)-1; p++ {
		if a[p] == -1 {
			continue
		}

		for i := p; i < len(a); i++ {
			if a[p] > a[i] && a[i] != -1 {
				a[p], a[i] = a[i], a[p]
			}

		}
	}

	return a
}
