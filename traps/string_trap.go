package main

import "fmt"

func main() {
	//Strings are essentially slices of bytes, and are not required to be UTF-8 encoded.
	// But the range loop assumes UTF-8 encoding. The iteration element is the
	// rune (like a character), and the index is the byte number of the first
	// byte of the rune (not the number of the rune). Thus, when the multi-byte
	// character é is encountered, the index 4 is skipped.
	str := "Café-society"

	for i, c := range str {
		fmt.Printf("Character #%d is %c \n", i, c)
	}
}