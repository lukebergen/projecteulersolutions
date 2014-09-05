package main

import "fmt"

func main() {
	sum, a, b, n := 0, 1, 0, 0
	for a < 4000000 {
		if a%2 == 0 {
			sum += a
		}
		n = a + b
		b = a
		a = n
	}

	fmt.Printf("Sum: %d", sum)
}
