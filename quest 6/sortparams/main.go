package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	length := len(args)
	if length == 1 {
		return
	} else if length == 2 {
		args = []string{args[1]}
	} else {
		args = args[1:]
		for i := 0; i < length-1; i++ {
			for j := 1; j < length-1; j++ {
				if args[j-1] > args[j] {
					args[j-1], args[j] = args[j], args[j-1]
				} else {
					continue
				}
			}
		}
	}

	for _, value := range args {

		for _, char := range value {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')

	}
}
