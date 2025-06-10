package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	var result string
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))  // Expected: "369\n"
	fmt.Print(ThirdTimeIsACharm(""))           // Expected: "\n"
	fmt.Print(ThirdTimeIsACharm("a b c d e f")) // Expected: "b e\n"
	fmt.Print(ThirdTimeIsACharm("12"))         // Expected: "\n"
}