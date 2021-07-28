import "regexp"

func variableName(name string) bool {
	regx := regexp.MustCompile("^\\d|\\s|[^a-zA-Z0-9_]")
	return !regx.Match([]byte(name))
}