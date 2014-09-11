package main

import (
	"fmt"
	"math"
)

func main() {
	c := 1
	var result int64
	for i := int64(3); true; i += 2 {
		if isPrime(i) {
			c++
		}

		if c == 10001 {
			result = i
			break
		}
	}
	fmt.Println("result: %d", result)
}

func isPrime(n int64) bool {
	start := int64(2)
	end := int64(math.Ceil(math.Sqrt(float64(n))))
	for i := start; i <= end; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
