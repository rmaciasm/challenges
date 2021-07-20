func arrayMaximalAdjacentDifference(inputArray []int) int {
	dif := -1000
	for i := 0; i < len(inputArray)-1; i++ {
		fmt.Println(inputArray[i], inputArray[i+1])
		if inputArray[i] == inputArray[i+1] {
			if dif < 0 {
				dif = 0
			}
			continue
		}

		if inputArray[i] > inputArray[i+1] {
			if inputArray[i]-inputArray[i+1] > dif {
				dif = inputArray[i] - inputArray[i+1]
			}
		} else {
			if inputArray[i+1]-inputArray[i] > dif {
				dif = inputArray[i+1] - inputArray[i]
			}

		}

	}
	return dif
}
