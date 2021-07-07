
func matrixElementsSum(matrix [][]int) int {
	sum := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				for m := r; m < len(matrix); m++ {
					matrix[m][c] = 0
				}
			} else {
				sum += matrix[r][c]
			}
		}
	}
	return sum
}
	
