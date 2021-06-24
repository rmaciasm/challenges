package main

func checkPalindrome(inputString string) bool {
	size := len(inputString)

	if size < 2 {
		return true
	}

	return inputString[0] == inputString[size-1] && checkPalindrome(inputString[1:size-1])

}
