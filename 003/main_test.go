package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestIsPrime(t *testing.T) {
	if !isPrime(3) {
		t.Error("3 should be prime")
	}
	if isPrime(4) {
		t.Error("4 should not be prime")
	}
	if !isPrime(17) {
		t.Error("17 should be prime")
	}
	if isPrime(24) {
		t.Error("24 should not be prime")
	}
}
