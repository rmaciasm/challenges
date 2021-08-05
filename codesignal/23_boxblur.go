
func boxBlur(image [][]int) (outputBox [][]int) {
	box, row, column := 9, len(image), len(image[0])
	outputBox = make([][]int, row-2)
	var sum, minRow, minColumn int
	for i := range outputBox {
		outputBox[i] = make([]int, column-2)
	}
	for i := 0; i < row-2; i++ {
		for j := 0; j < column-2; j++ {
			for k := i; minRow < 3; k++ {
				for l := j; minColumn < 3; l++ {
					sum += image[k][l]
					minColumn++
				}
				minColumn = 0
				minRow++
			}
			sum /= box
			outputBox[i][j] = sum
			minRow, sum = 0, 0
		}
	}
	return
}

	
