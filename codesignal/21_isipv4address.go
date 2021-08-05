import "regexp"
import "net"

func isIPv4Address(inputString string) bool {
	isValid := regexp.MustCompile(`^([^0]\d{0,2}|0)(?:\.([^0]\d{0,2}|0)){3}$`)
	if !isValid.MatchString(inputString) {
		return false
	}
	return net.ParseIP(inputString) != nil
}
