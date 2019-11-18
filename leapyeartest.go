package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputYear = os.Args[1]
	var year, error = strconv.Atoi(inputYear)
	if error != nil {
		fmt.Println("Failed to convert ")
	}
	//fmt.Printf("Errors were %v", error)
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%v is a leap year \n", year)
	} else {
		fmt.Println("Try again")
	}

}
