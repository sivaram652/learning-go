package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if num := rand.Intn(2); num == 0 {
		fmt.Println("Number is 0")
	} else {
		fmt.Println("Number is not 0")
	}

}
