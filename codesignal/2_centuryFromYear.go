package main

import "fmt"

func centuryFromYear(year int) int {
	res := year % 100
	total := (year - res) / 100
	if res != 0 {
		total += 1
	}
	return total
}

func main() {
	fmt.Println("💜", centuryFromYear(1905))
}
