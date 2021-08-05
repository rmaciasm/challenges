	
func minesweeper(matrix [][]bool) (newMatrix [][]int) {
	lR, lC := len(matrix), len(matrix[0])
	newMatrix = make([][]int, lR)
	for k, _ := range newMatrix {
		newMatrix[k] = make([]int, lC)
	}
	for rk, rv := range matrix {
		for ck, _ := range rv {
			isUp := rk-1 >= 0
			isDown := rk+1 < lR
			isLeft := ck-1 >= 0
			isRight := ck+1 < lC

			if isUp {
				if matrix[rk-1][ck] {
					newMatrix[rk][ck] += 1
				}

				if isRight {
					if matrix[rk-1][ck+1] {
						newMatrix[rk][ck] += 1
					}
				}

				if isLeft {
					if matrix[rk-1][ck-1] {
						newMatrix[rk][ck] += 1
					}
				}

			}

			if isDown {
				if matrix[rk+1][ck] {
					newMatrix[rk][ck] += 1
				}

				if isRight {
					if matrix[rk+1][ck+1] {
						newMatrix[rk][ck] += 1
					}
				}

				if isLeft {
					if matrix[rk+1][ck-1] {
						newMatrix[rk][ck] += 1
					}
				}
			}

			if isRight {
				if matrix[rk][ck+1] {
					newMatrix[rk][ck] += 1
				}
			}

			if isLeft {
				if matrix[rk][ck-1] {
					newMatrix[rk][ck] += 1
				}
			}
		}

	}

	return

}

	
