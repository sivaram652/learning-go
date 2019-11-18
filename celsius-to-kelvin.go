package main

import (
	"fmt"
	"os"
	"strconv"
)

type kelvin float64
type celsius float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	c := celsius(f)
	fmt.Printf("%v\n", celsiusToKelvin(c))
}
