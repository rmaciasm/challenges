func depositProfit(deposit int, rate int, threshold int) int {
	years := 0
	quantity := deposit

	for quantity < threshold {
		quantity += quantity * rate / 100
		years++
	}

	return years
}
