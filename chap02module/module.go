package chap02module

import (
	"encoding/base64"
	"fmt"
)

func Hello() {
	fmt.Println("hello")
}

func StringTobase64(inString string) {
	encodedString := base64.StdEncoding.EncodeToString([]byte(inString))
	fmt.Printf("Original string: %s\n", inString)
	fmt.Printf("After base64 encoding: %s\n", encodedString)
}
