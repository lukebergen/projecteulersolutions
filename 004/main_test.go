package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if isPalindrome(1234) {
		t.Error("1234 should not have been a palindrome")
	}
	if isPalindrome(12345) {
		t.Error("12345 should not have been a palindrome")
	}
	if !isPalindrome(1221) {
		t.Error("1221 should not have been a palindrome")
	}
	if !isPalindrome(12321) {
		t.Error("12321 should have been a palindrome")
	}
}

// 71373532 ns/op
func BenchmarkMainSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainSync()
	}
}

// 74767508 ns/op
func BenchmarkMainAsync1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(1)
	}
}

// 74928701 ns/op
func BenchmarkMainAsync2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(2)
	}
}

// 73607226 ns/op
func BenchmarkMainAsync10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(10)
	}
}

// 74910953 ns/op
func BenchmarkMainAsync20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(20)
	}
}
