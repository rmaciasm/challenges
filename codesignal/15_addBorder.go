import (
	"strings"
)

func addBorder(picture []string) []string {
	size := len(picture)
	withBorder := make([]string, size+2)
	lenght := len(picture[0])
	word := make([]string, 0)
	for k := 0; k < lenght+2; k++ {
		word = append(word, "*")
	}
	fword := strings.Join(word, "")
	withBorder[0] = fword

	for i := 0; i < size; i++ {
		withBorder[i+1] = "*" + picture[i] + "*"
	}

	withBorder[size+1] = fword

	return withBorder
}

