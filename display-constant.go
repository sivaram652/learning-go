package main

import (
	"fmt"
)

func main() {
	const distance = 236000000000000000
	const oneLightYear = 299792 * 86400 * 365

	fmt.Println("One light year in KM =  ", oneLightYear, " km")
	fmt.Println("Canis Majoris is ", distance/oneLightYear, " light years away")
}
