package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsCapitalized(s string) bool {
	// Return false if string is empty
	if s == "" {
		return false
	}
	
	// Split the string into words
	words := strings.Fields(s)
	
	// Check each word
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		
		// Get the first character of the word
		firstChar := rune(word[0])
		
		// If it's a letter and lowercase, return false
		if unicode.IsLetter(firstChar) && unicode.IsLower(firstChar) {
			return false
		}
	}
	
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))  // false
	fmt.Println(IsCapitalized("Hello How Are You"))    // true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))    // true
	fmt.Println(IsCapitalized("Whatsthis4"))           // true
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))       // true
	fmt.Println(IsCapitalized(""))                     // false
}