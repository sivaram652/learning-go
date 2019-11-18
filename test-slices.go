package main

import (
	"fmt"
)

func main() {
	planets := [...]string{"Mer", "Ven", "Ear", "Mar"}
	outerPlanets := [...]string{"Jup", "Sat", "Ura", "Nep", "Plu"}
	inner := planets[0:3]
	secondInner := inner
	inner = outerPlanets[0:2]

	fmt.Printf("inner %v\n", inner)
	secondInner[2] = "NewEar"
	fmt.Printf("2nd inner %v\n", secondInner)
	fmt.Printf("Planets %v\n", planets)
	fmt.Printf("planets type - %T, inner planets type - %T\n", planets, inner)
}
