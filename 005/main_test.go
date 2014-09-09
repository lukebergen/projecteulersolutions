package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestDivisible(t *testing.T) {
	if divisible(17) {
		t.Error("17 should not be 'divisible'")
	}
	bigCommonMultiple := int64(1)
	for i := int64(2); i <= 20; i++ {
		bigCommonMultiple *= i
	}
	if !divisible(bigCommonMultiple) {
		t.Error("%d should be 'divisible'", bigCommonMultiple)
	}
}
