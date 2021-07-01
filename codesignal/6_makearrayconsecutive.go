func makeArrayConsecutive2(statues []int) (total int) {
	sort.Ints(statues)

	for k := 0; k < len(statues)-1; k++ {

		if r := statues[k+1] - statues[k]; r > 1 {
			total += (r - 1)
		}
	}

	return
}
