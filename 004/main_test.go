package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if x := isPalindrome(1234); x {
		t.Error("1234 should not have been a palindrome")
	}
	if x := isPalindrome(12345); x {
		t.Error("12345 should not have been a palindrome")
	}
	if x := isPalindrome(1221); !x {
		t.Error("1221 should not have been a palindrome")
	}
	if x := isPalindrome(12321); !x {
		t.Error("12321 should have been a palindrome")
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
