package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x = os.Args[1]
	var y, err = strconv.Atoi(x)
	if err == nil {
		switch y {
		case 1:
			fmt.Println("I am one")
		case 2:
			fmt.Println("I am two")
			fallthrough
		case 3:
			fmt.Println("I am also a prime number")
		}
	}
}
