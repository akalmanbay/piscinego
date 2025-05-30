package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	var arg []rune

	if len(arguments) > 1 {
		arg = []rune(arguments[1])
	} else {
		arg = []rune(arguments[0])
	}
	for i, value := range arg {
		if i > 1 {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}
