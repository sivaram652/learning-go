package main

import (
	"fmt"
)

func main() {
	var red uint16 = 65535
	//red++
	red++
	fmt.Println(red)
	var number int8 = -128
	//	number++
	number--
	fmt.Println(number)

	fmt.Printf("%b\n", number)
}
