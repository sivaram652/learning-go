package main

import (
	"fmt"
)

func main() {
	whole := 45
	real := 36.45
	yes := true
	str := "hello"

	fmt.Printf("%v is of Type %[1]T\n", whole)
	fmt.Printf("%v is of Type %[1]T\n", real)
	fmt.Printf("%v is of Type %[1]T\n", yes)
	fmt.Printf("%v is of Type %[1]T\n", str)
}
