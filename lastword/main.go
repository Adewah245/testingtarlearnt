package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]

	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	if end < 0 {
		return
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	start++

	fmt.Println(s[start : end+1])
}