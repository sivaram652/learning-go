package main

import (
	"fmt"
)

func main() {
	message := "a L fdph, L vdz, L frqtxhuhg."
	//var decodedMessage string

	for i := 0; i < len(message); i++ {
		c := message[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			c = c - 3

			if c < 'a' || c < 'A' {
				c = c + 26
			}
		}
		fmt.Printf("%c, %v\n", c, c)
		//fmt.Printf("%c", c)
	}
}
