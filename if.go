package main

import "fmt"

func main() {
	var inputAge int8 = 13

	if inputAge == 13 {
		fmt.Println("Person is starting teen")
	} else if inputAge < 18 {
		fmt.Println("Person is minor")
	} else {
		fmt.Println("Person is major")
	}
}
