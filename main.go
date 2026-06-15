package main

import (
	"fmt"
	"strings"
)

// IsPalindrome checks if a string is the same forward and backward
func IsPalindrome(word string) bool {
	// Convert to lowercase to handle capitalization
	word = strings.ToLower(word)
	
	shortLen := len(word)
	for i := 0; i < shortLen/2; i++ {
		if word[i] != word[shortLen-1-i] {
			return false // Characters don't match
		}
	}
	return true // It's a palindrome
}

func main() {
	testWord := "Radar"
	result := IsPalindrome(testWord)
	
	fmt.Printf("Is '%s' a palindrome? %t\n", testWord, result)
}

