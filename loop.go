package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var degrees int = 0
	var count int = 0
	for {
		fmt.Println(degrees)
		degrees++
		count++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(10) == 0 {
				break
			}
		}
		fmt.Printf("Count = %v", count)
	}
}
