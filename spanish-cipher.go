package main

import (
	"fmt"
)

func main() {

	spanishMessage := "Hola Estación Espacial Internacional"
	for _, c := range spanishMessage {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else {
			if c >= 'A' && c <= 'Z' {
				c = c + 13
				if c > 'Z' {
					c = c - 26
				}
			} else {
				if string(c) != " " {
					c = c + 13
				}
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
