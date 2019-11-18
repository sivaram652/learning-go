package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var aNickel uint16 = 5
	var aDime uint16 = 10
	var aQuarter uint16 = 25
	denominations := []uint16{aNickel, aDime, aQuarter}

	var piggyMinBalance uint16 = 2000
	var currPiggyBalance uint16 = 0
	var denomination uint16
	var currPiggyBalanceInDecimals float32

	for piggyMinBalance-currPiggyBalance > 0 {
		denomination = denominations[rand.Intn(3)]
		fmt.Println("Adding denomination : ", denomination)
		currPiggyBalance += denomination
		currPiggyBalanceInDecimals = float32(currPiggyBalance) / float32(100)
		fmt.Printf("Current Balance : %5.2f\n", currPiggyBalanceInDecimals)
		fmt.Printf("Type %T for %v\n", currPiggyBalance, currPiggyBalance)
		fmt.Println("--------------------------")
	}
}
