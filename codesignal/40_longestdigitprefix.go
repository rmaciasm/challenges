import "regexp"

func longestDigitsPrefix(inputString string) string {
	re := regexp.MustCompile(`^(\d*)`)
	val := re.Find([]byte(inputString))
	return string(val)
}
