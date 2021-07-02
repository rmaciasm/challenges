
func almostIncreasingSequence(sequence []int) bool {
	sequenceCopy := make([]int, len(sequence))
	copy(sequenceCopy, sequence)
	i := findIsIncreasingSequence(sequence)
	if i == -1 {
		return true
	}
	if findIsIncreasingSequence(append(sequence[:i], sequence[i+1:]...)) == -1 {
		return true
	}
	if findIsIncreasingSequence(append(sequenceCopy[:i+1], sequenceCopy[i+2:]...)) == -1 {
		return true
	}
	return false

}

func findIsIncreasingSequence(sequence []int) int {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			return i
		}
	}
	return -1
}

	
