package main

import (
	"fmt"
	"os"
)

func main() {
	var room = os.Args[1]

	if room == "cave" {
		fmt.Println("This is a cave")
	} else if room == "entrance" {
		fmt.Println("This is an entrance")
	} else if room == "mountain" {
		fmt.Println("This is a mountain")
	} else {
		fmt.Println("We don't have such a room")
	}

}
