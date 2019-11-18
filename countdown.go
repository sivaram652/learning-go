package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 100
	for count > 0 {
		time.Sleep(time.Second)
		if rand.Intn(99) == 0 {
			fmt.Println("Count is ", count)

			break
		}
		count--
	}

}
