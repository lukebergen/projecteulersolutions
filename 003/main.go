package main

import (
	"fmt"
	"math"
)

const n int64 = 600851475143

func main() {
	primes := []int64{2, 3, 5}
	result := n
	for {
		if isPrime(result) {
			break
		}
		foundPrime := false
		for i := 0; i < len(primes); i++ {
			if result%primes[i] == 0 {
				result = result / primes[i]
				foundPrime = true
				break
			}
		}
		if !foundPrime {
			if !isPrime(result) {
				nextPrime := primes[len(primes)-1] + 1
				for ; !isPrime(nextPrime); nextPrime++ {
				}
				primes = append(primes, nextPrime)
			}
		}
	}
	fmt.Printf("result: %d\n", result)
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
