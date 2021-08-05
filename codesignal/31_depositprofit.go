func depositProfit(deposit int, rate int, threshold int) int {
	years := 0
	quantity := float64(deposit)

	for quantity < float64(threshold) {
		quantity += quantity * float64(rate) / 100.
		years++
	}

	return years
}

