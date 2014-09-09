package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestPow(t *testing.T) {
	if pow(2, 3) != 8 {
		t.Error("pow(2, 3) should have been 8")
	}
}
