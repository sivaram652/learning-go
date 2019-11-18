package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var longestNameLength int
	spacelines := []string{"Space X", "Virgin Galactic", "Space Adventures"}
	tripTypes := []string{"One-Way", "RoundTrip"}
	const distanceToMarsInKm int = 62100000
	const tripTypeOneWay string = "One-Way"
	const tripTypeReturn string = "Round-trip"
	for i := 0; i < len(spacelines); i++ {
		if i == 0 {
			longestNameLength = len(spacelines[i])
		} else {
			if longestNameLength < len(spacelines[i]) {
				longestNameLength = len(spacelines[i])
			}
		}
	}

	fmt.Printf("%-17v", "Spaceline")
	fmt.Print("Days ")
	fmt.Print("Trip Type  ")
	fmt.Print("Price ")
	fmt.Print("Speed\n")
	fmt.Println("=======================================")

	for count := 10; count > 0; count-- {

		spaceline := spacelines[rand.Intn(3)]
		speedInKmPerSec := rand.Intn(14) + 16
		//var priceInMillionDollars int

		priceInMillionDollars := speedInKmPerSec + 20
		oneWayTripDurationInDays := distanceToMarsInKm / (speedInKmPerSec * 24 * 3600)
		fmt.Printf("%-17v", spaceline)
		fmt.Printf("%4v ", oneWayTripDurationInDays)
		tripType := tripTypes[rand.Intn(2)]
		fmt.Printf("%-12v", tripType)
		if tripType == "RoundTrip" {
			fmt.Printf("%4v", priceInMillionDollars*2)
		} else {
			fmt.Printf("%4v", priceInMillionDollars)
		}

		fmt.Printf("%5v\n", speedInKmPerSec)
	}
	//fmt.Printf("Spaceline = %v, Speed = %v, Duration = %v", spaceline, speedInKmPerSec, oneWayTripDurationInDays)

	//}
	//const spaclineSpaceX string = "Space X"
	//const spacelineVirginGalc string = "Virgin Galactic"
	//const spacelineSpaceAdv string = "Space Adventures"
	/*
		fmt.Printf("%-10v\n", 25)
		fmt.Printf("%6v\n", 34)
		fmt.Println("Length of array is ", len(spacelines))
		fmt.Println("Longest spaceline ", longestNameLength)
	*/
	//	for count := 100; count > 0; count-- {
	//		fmt.Println(rand.Intn(14) + 16)
	//	}

}
