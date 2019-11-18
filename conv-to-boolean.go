package main

import (
	"fmt"
	"os"
)

func main() {
	var command = os.Args[1]
	var b bool
	switch command {
	case "yes", "true", "1":
		b = true
		fmt.Printf("You entered %v\n", b)
	case "no", "false", "0":
		b = false
		fmt.Printf("You entered %v\n", b)
	default:
		fmt.Println("You have entered an invalid value.")
	}

}
