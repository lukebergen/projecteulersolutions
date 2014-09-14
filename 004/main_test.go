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

func BenchmarkMainSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainSync()
	}
}

func BenchmarkMainAsync1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(1)
	}
}

func BenchmarkMainAsync2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(2)
	}
}

func BenchmarkMainAsync10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(10)
	}
}

func BenchmarkMainAsync20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainAsync(20)
	}
}
