package main

import (
	"fmt"
)

func main() {
	var a int8 = 62
	var b int8 = 65
	var mod int8 = a % b % 26
	fmt.Println("Modulus - ", mod)
	var c int8
	c = mod + 39
	fmt.Printf("Char is %c\n", rune(c))

}
