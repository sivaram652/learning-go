package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	for count := 10; count > 0; count-- {
		year := rand.Intn(100) + 2000
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		fmt.Println(era, year, month, daysInMonth)
	}

}
