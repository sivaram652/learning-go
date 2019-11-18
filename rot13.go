package main

import (
	"fmt"
)

func main() {
	var decipheredMessage string
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {

			fmt.Printf("Before adding 13, c is %c\n", c)
			c = c + 13

			fmt.Printf("After adding 13, c is %c\n", c)
			if c > 'z' {
				c = c - 26
				fmt.Printf("After subtracting 26, c is %c\n", c)
			}
		}
		fmt.Println("**********************")
		decipheredMessage = decipheredMessage + string(c)
	}

	fmt.Printf("Final messge is %v", decipheredMessage)

}
