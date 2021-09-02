import "strings"

func findEmailDomain(address string) string {
	parts := strings.Split(address, "@")
	return parts[len(parts)-1]
}