package main

import "fmt"

func HashCode(dec string) string {
    result := ""
    strLen := len(dec)
    
    for _, char := range dec {
        // Get ASCII value of current character
        ascii := int(char)
        
        // Apply hash equation: (ASCII + string length) % 127
        hashedValue := (ascii + strLen) % 127
        
        // If result is unprintable (< 33), add 33
        if hashedValue < 33 {
            hashedValue += 33
        }
        
        // Convert back to character and append to result
        result += string(rune(hashedValue))
    }
    
    return result
}

func main() {
    fmt.Println(HashCode("A"))           // Expected: B
    fmt.Println(HashCode("AB"))          // Expected: BD
    fmt.Println(HashCode("BAC"))         // Expected: EDF
    fmt.Println(HashCode("Hello World")) // Expected: Spwwz+bz}wo
}