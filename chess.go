package main

import (
	"fmt"
)

func main() {
	var pieces [8][8]rune
	pieces[0] = [8]rune{'r', 'k', 'b', 'q', 'k', 'b', 'k', 'r'}
	for i := range pieces {
		pieces[1][i] = 'p'
		pieces[6][i] = 'P'
	}
	pieces[7] = [8]rune{'R', 'K', 'B', 'Q', 'K', 'B', 'K', 'R'}

	//fmt.Printf("%c", pieces)

	for i := range pieces {
		for j := range pieces[i] {
			fmt.Printf("%c ", pieces[i][j])
		}
		fmt.Printf("\n")
	}
}
