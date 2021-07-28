func chessBoardCellColor(cell1 string, cell2 string) bool {
	coord1 := []rune(cell1)
	coord2 := []rune(cell2)

	sum1 := coord1[0] + coord1[1]
	sum2 := coord2[0] + coord2[1]
	return sum1%2 == sum2%2
}
