package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	fmt.Printf("Time taken to reaach Andromeda is %v seconds\n", seconds)

	days := new(big.Int)

	days.Div(seconds, secondsPerDay)
	fmt.Printf("It will take %v days to reach Andromeda\n ", days)

}
