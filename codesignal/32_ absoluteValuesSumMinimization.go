
import (
	"math"
)

func absoluteValuesSumMinimization(a []int) int {
	smallest := math.MaxInt64
	number := 0

	for _, x := range a {
		sum := 0.0
		for _, v := range a {
			sum += math.Abs(float64(v) - float64(x))
		}

		if sum < float64(smallest) {
			smallest = int(sum)
			number = x
		}

	}

	return number

}
