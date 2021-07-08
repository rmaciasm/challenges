

import (
	"strings"
)

func commonCharacterCount(s1 string, s2 string) (count int) {
	if len(s1) >= len(s2) {
		count = returnCharCount(s1, s2)
	} else {
		count = returnCharCount(s2, s1)
	}

	return count

}

func returnCharCount(bigger string, smaller string) (count int) {
	for _, v := range bigger {
		if i := strings.Index(smaller, string(v)); i != -1 {
			small := strings.Split(smaller, "")
			small[i] = "1"
			smaller = strings.Join(small, "")
			count++
		}
	}
	return
}

