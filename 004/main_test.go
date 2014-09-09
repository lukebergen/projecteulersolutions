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

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
