package main

import (
	"fmt"
	"os"
)

func main() {
	var firstArg string = os.Args[0]
	var secondArg string = os.Args[1]

	fmt.Printf("First arg is %v. Second arg is %v", firstArg, secondArg)
}
