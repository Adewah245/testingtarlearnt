package main

import "fmt"

func Split(s, sep string) []string {
    if sep == "" {
        // If separator is empty, return slice with original string
        return []string{s}
    }
    
    if s == "" {
        // If string is empty, return slice with empty string
        return []string{""}
    }
    
    var result []string
    sepLen := len(sep)
    start := 0
    
    for i := 0; i <= len(s)-sepLen; i++ {
        // Check if we found the separator at position i
        if s[i:i+sepLen] == sep {
            // Add the substring from start to current position
            result = append(result, s[start:i])
            // Move start position past the separator
            start = i + sepLen
            // Skip ahead to avoid overlapping matches
            i += sepLen - 1
        }
    }
    
    // Add the remaining part of the string
    result = append(result, s[start:])
    
    return result
}

func main() {
    s := "HelloHAhowHAareHAyou?"
    fmt.Printf("%#v\n", Split(s, "HA"))
    
    // Additional test cases
    fmt.Printf("%#v\n", Split("a,b,c", ","))
    fmt.Printf("%#v\n", Split("hello", ""))
    fmt.Printf("%#v\n", Split("", "x"))
    fmt.Printf("%#v\n", Split("aaabaaab", "aa"))
}