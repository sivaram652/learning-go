package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	var keywordCharPosition int8
	var decipheredChar rune
	var message string
	cipherTextLength := len(cipherText)
	fmt.Println("cipher text length", cipherTextLength)
	keywordLength := len(keyword)
	fmt.Println("keyword length", keywordLength)
	keywordRepeats := (cipherTextLength / keywordLength) + 1
	fmt.Println("keywordRepeats ", keywordRepeats)
	newKeyword := strings.Repeat(keyword, keywordRepeats)
	fmt.Println("newKeyword ", newKeyword)
	fmt.Println("length of keyword ", keywordLength)
	fmt.Println(" ######## ", cipherText)
	fmt.Println(" ######## ", newKeyword)
	for i, c := range cipherText {
		/*c = c - 65
		keywordCharPosition = int8(newKeyword[i]) - 65
		decipheredChar = (c-rune(keywordCharPosition)+26)%26 + 65*/
		keywordCharPosition = int8(newKeyword[i]) - 65
		if (c - rune(keywordCharPosition)) >= 65 {
			decipheredChar = c - rune(keywordCharPosition)
		} else {
			decipheredChar = c - rune(keywordCharPosition) + rune(26)
		}

		fmt.Printf("Deciphered Char %c, %v\n ", decipheredChar, decipheredChar)
		message = message + string(decipheredChar)

		//fmt.Printf("Original %c - %v, Keyword %c - %v, New %c - %v\n", c, c, newKeyword[i], keywordCharPosition)
	}
	fmt.Printf("%v", message)
}
