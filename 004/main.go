package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Printf("result: %d\n", mainAsync(1)) //    72818811 ns/op
	// fmt.Printf("result: %d\n", mainAsync(2)) //    72818811 ns/op
	// fmt.Printf("result: %d\n", mainAsync(10)) //   73350902 ns/op
	// fmt.Printf("result: %d\n", mainAsync(20)) //   73034246 ns/op
	fmt.Printf("result: %d\n", mainSync()) //      70225353 ns/op
}

func mainAsync(threads int) int {
	palindromes := make(chan int, threads)
	done := make(chan bool)

	for i := 0; i < threads; i++ {
		go searchForPalindromes(palindromes, done, i, threads)
	}

	best := 0
	n := 0
	gors := threads
	for gors > 0 {
		select {
		case n = <-palindromes:
			// fmt.Printf("Received palindrome %d off channel\n", n)
			if n > best {
				best = n
			}
		case <-done:
			gors--
		}
	}
	return best
}

func searchForPalindromes(palindromes chan int, done chan bool, start int, jump int) {
	for i := 100 + start; i < 1000; i += jump {
		for j := 100; j < 1000; j++ {
			if isPalindrome(i * j) {
				// fmt.Printf("Found palindrome: %d. Pushing it onto channel\n", i*j)
				palindromes <- i * j
			}
		}
	}
	// fmt.Printf("Done searching my section of the numbers. Sending -1")
	done <- true
}

func mainSync() int {
	best := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrome(i*j) && i*j > best {
				best = i * j
			}
		}
	}
	return best
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
