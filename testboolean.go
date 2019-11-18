package main

import (
	"fmt"
	"strings"
)

func main() {
	var wearShades bool
	var caveSign string = "Did you read the warning"
	fmt.Println("Wear shades is", wearShades)
	fmt.Println("The warning sign contains read:", strings.Contains(caveSign, "read"))
}
