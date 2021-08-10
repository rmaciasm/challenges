func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	grow, day := 0, 0

	for {
		day++
		grow += upSpeed
		if grow >= desiredHeight {
			return day
		}

		grow -= downSpeed
	}
}
