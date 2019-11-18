package main

import (
	"fmt"
	"os"
	"strconv"
)

func kelvinToCelsius(k float64) float64 {
	c := k - 273.15
	return c
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32
	return f
}

func main() {
	k, _ := strconv.ParseFloat(os.Args[1], 64)
	c, _ := strconv.ParseFloat(os.Args[2], 64)
	k2, _ := strconv.ParseFloat(os.Args[3], 64)
	celsius := kelvinToCelsius(k)
	farenheit := celsiusToFahrenheit(c)

	f := celsiusToFahrenheit(kelvinToCelsius(k2))

	fmt.Printf(" Celsius = %v, Fahrenheit = %v, Fahrenheit2 = %v\n", celsius, farenheit, f)
}
