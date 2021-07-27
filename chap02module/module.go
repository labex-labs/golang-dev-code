package hello

import (
	"encoding/base64"
	"fmt"
)

// 输出 hello
func Hello() {
	fmt.Println("hello")
}

// 将输入的字符串转为 base64 编码
func StringTobase64(inString string) {
	// 编码为 base64
	encodedString := base64.StdEncoding.EncodeToString([]byte(inString))
	fmt.Printf("原字符串: %s\n", inString)
	fmt.Printf("base64 编码后: %s\n", encodedString)
}
