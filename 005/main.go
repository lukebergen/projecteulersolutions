package main

import "fmt"

func main() {
	var result int64 = 20
	for ; !divisible(result); result += 20 {
	}
	fmt.Println("result: %d", result)
}

func divisible(n int64) bool {
	divisors := []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	l := len(divisors)

	for i := 0; i < l; i++ {
		if n%divisors[i] != 0 {
			return false
		}
	}
	return true
}
