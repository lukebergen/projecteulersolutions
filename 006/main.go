package main

import "fmt"

func main() {
	sumOfSquares := int64(1)
	sumToOneHundred := 1
	for i := 2; i <= 100; i++ {
		sumOfSquares += pow(i, 2)
		sumToOneHundred += i
	}

	squareOfSums := pow(sumToOneHundred, 2)

	fmt.Println("result: %d", squareOfSums-sumOfSquares)
}

func pow(x, y int) (result int64) {
	result = int64(x)
	xx := int64(x)
	for i := 1; i < y; i++ {
		result *= xx
	}
	return
}
