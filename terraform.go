package main

import (
	"fmt"
)

func terraform(planets [4]string) [4]string {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}

	return planets
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
	}

	planets = terraform(planets)
	fmt.Println(planets)
}
