package main

import (
	"fmt"
	"os"
	"strconv"
)

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) farenheit() farenheit {
	return k.celsius().farenheit()
}

// conversions from farenheit
type farenheit float64

func (f farenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (f farenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

//conversions from celsius
type celsius float64

func (c celsius) farenheit() farenheit {
	return farenheit(9/5*c + 32)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {
	t := os.Args[1]
	units := os.Args[2]
	convertTo := os.Args[3]
	// source temp
	var k kelvin
	var c celsius
	var f farenheit

	// target temp
	var cel celsius
	var kel kelvin
	var far farenheit

	var flOut float64

	switch units {
	case "K":
		fl, _ := strconv.ParseFloat(t, 64)
		k = kelvin(fl)
		switch convertTo {
		case "C":
			cel = k.celsius()
			flOut = float64(cel)
		case "F":
			far = k.farenheit()
			flOut = float64(far)
		}
	case "C":
		fl, _ := strconv.ParseFloat(t, 64)
		c = celsius(fl)
		switch convertTo {
		case "K":
			kel = c.kelvin()
			flOut = float64(kel)
		case "F":
			far = c.farenheit()
			flOut = float64(far)
		}
	case "F":
		fl, _ := strconv.ParseFloat(t, 64)
		f = farenheit(fl)
		switch convertTo {
		case "C":
			cel = f.celsius()
			flOut = float64(cel)
		case "K":
			kel = f.kelvin()
			flOut = float64(kel)
		}
	}

	fmt.Println(" Temp id ", flOut)

}
