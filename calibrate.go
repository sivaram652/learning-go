package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
func main() {

	var offset kelvin = 5.0
	s := calibrate(realSensor, offset)
	fmt.Println(s())
	offset = 7.5
	s = calibrate(realSensor, offset)
	fmt.Println(s())
	fmt.Println("---------")
	s2 := calibrate(fakeSensor, offset)
	fmt.Println(s2())
	fmt.Println(s2())
	fmt.Println(s2())
	fmt.Println(s2())

}
