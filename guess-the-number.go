package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var numberString = os.Args[1]
	var numberToGuess, err = strconv.Atoi(numberString)
	var computerGuess int
	if err == nil {
		for {
			computerGuess = rand.Intn(100)
			fmt.Printf("Computer's guess was = %v \n", computerGuess)
			fmt.Println("Number to guess < Computer Guess : ", (numberToGuess < computerGuess))

			if computerGuess == numberToGuess {
				break
			}
		}

		fmt.Println("Number to guess was = ", computerGuess)
	}
}
