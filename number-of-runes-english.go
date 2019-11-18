package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	count := utf8.RuneCountInString(alphabets)
	//_, bytecount := utf8.DecodeRuneInString(alphabets)
	spanishq := "¿"

	fmt.Printf("Number of runes is %v and Number of bytes is %v and Number of bytes in ¿ is %v", count, len(alphabets), len(spanishq))

}
