package main

import (
	"fmt"
	"os"
)

func main() {
	var command = os.Args[1]
	switch command {
	case "enter cave":
		fmt.Println("Entered the cave")
	case "go south":
		fmt.Println("Going far far away")
	default:
		fmt.Println("Going to the Nether-lands")
	}
}
