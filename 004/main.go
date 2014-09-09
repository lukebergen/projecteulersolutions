package main

import (
	"fmt"
	"strconv"
)

func main() {
	best := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if i*j > best && isPalindrome(i*j) {
				best = i * j
			}
		}
	}
	fmt.Println("result: %d", best)
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	l := len(str) - 1
	for i := 0; i < l/2+1; i++ {
		if str[i] != str[l-i] {
			return false
		}
	}
	return true
}
