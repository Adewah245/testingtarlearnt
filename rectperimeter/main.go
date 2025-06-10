package main

import "fmt"

// RectPerimeter calculates the perimeter of a rectangle
// Returns -1 if either width or height is negative
func RectPerimeter(w, h int) int {
	// Check if either argument is negative
	if w < 0 || h < 0 {
		return -1
	}
	
	// Calculate perimeter: 2 * (width + height)
	return 2 * (w + h)
}

func main() {
	fmt.Println(RectPerimeter(10, 2))        // Expected: 24
	fmt.Println(RectPerimeter(434343, 898989)) // Expected: 2666664
	fmt.Println(RectPerimeter(10, -2))       // Expected: -1
}