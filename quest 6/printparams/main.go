package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i := range arguments {
		if i != 0 {
			arg := []rune(arguments[i])
			for j := range arg {
				z01.PrintRune(arg[j])
			}
			z01.PrintRune('\n')
		}
	}
}
