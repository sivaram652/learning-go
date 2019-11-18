package main

import (
	"fmt"
)

func main() {
	var color uint8 = 25
	fmt.Printf("hex is %x\n", color)

	var num uint8 = 255
	num++
	num++
	num++
	num++
	fmt.Printf("num is %v, in binary it is %08[1]b\n", num)
}
