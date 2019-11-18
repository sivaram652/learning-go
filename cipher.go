package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	var m rune
	var message string
	plainText = strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	newKeyword := strings.Repeat(keyword, (len(plainText)/len(keyword))+1)

	fmt.Printf("%v %v\n", plainText, newKeyword)

	var k int8
	for i, c := range plainText {
		fmt.Printf("c BEFORE %c %v\n", c, c)
		c = c - 'A'
		fmt.Printf("c AFTER %c %v\n", c, c)
		k = int8(newKeyword[i])
		fmt.Printf("k BEFORE %c %v\n", k, k)
		k = int8(newKeyword[i]) - int8('A')
		fmt.Printf("k AFTER %c %v\n", k, k)
		fmt.Println("====================")

		m = (c+rune(k))%26 + 'A'
		fmt.Printf("%c\n", m)
		message = message + string(m)
	}
	fmt.Println(message)

	// Now decipher the above message
	var keywordCharPosition int8
	var decipheredChar rune
	var cipheredMessage string
	for i, d := range message {
		keywordCharPosition = int8(newKeyword[i]) - 65
		if (d - rune(keywordCharPosition)) >= 65 {
			decipheredChar = d - rune(keywordCharPosition)
		} else {
			decipheredChar = d - rune(keywordCharPosition) + rune(26)
		}

		fmt.Printf("Deciphered Char %c, %v\n ", decipheredChar, decipheredChar)
		cipheredMessage = cipheredMessage + string(decipheredChar)

	}
	fmt.Println(cipheredMessage)
}
