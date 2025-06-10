package main

import "fmt"

func FishAndChips(n int) string {
	// Check if number is negative
	if n < 0 {
		return "error: number is negative"
	}
	
	// Check if divisible by both 2 and 3 (i.e., divisible by 6)
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	
	// Check if divisible by 2
	if n%2 == 0 {
		return "fish"
	}
	
	// Check if divisible by 3
	if n%3 == 0 {
		return "chips"
	}
	
	// If not divisible by 2 or 3
	return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(4))  // fish
	fmt.Println(FishAndChips(9))  // chips
	fmt.Println(FishAndChips(6))  // fish and chips
	
	// Additional test cases
	fmt.Println(FishAndChips(-5)) // error: number is negative
	fmt.Println(FishAndChips(7))  // error: non divisible
	fmt.Println(FishAndChips(0))  // fish and chips (0 is divisible by both 2 and 3)
}