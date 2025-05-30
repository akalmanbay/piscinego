package main

import (
	"os"

	"github.com/01-edu/z01"
)

func StringToPrint(s string) {
	for _, char := range s {
		z01.PrintRune(rune(char))
	}
}

func RunesToPrint(s []rune) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func Insert(s1 string, s2 string) string {
	return s1 + s2
}

func Order(s string) []rune {
	res := []rune(s)
	for range s {
		for i := 0; i < len(s)-1; i++ {
			if res[i] > res[i+1] {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
	return res
}

func main() {
	args := os.Args
	length := len(args)
	res := args[length-1]
	runes := []rune{}
	helpText := "--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument."

	isOrder := false
	isHelp := false

	for _, arg := range args {
		if len(arg) > 7 && arg[:8] == "--insert" {
			res = Insert(res, arg[9:])
		} else if len(arg) > 1 && arg[:2] == "-i" {
			res = Insert(res, arg[3:])
		} else if arg == "--order" || arg == "-o" {
			runes = Order(res)
			isOrder = true
		} else if arg == "--help" || arg == "-h" {
			isHelp = true
		}
	}
	if isHelp || length == 1 {
		StringToPrint(helpText)
	} else if isOrder {
		RunesToPrint(runes)
	} else {
		StringToPrint(res)
	}
	z01.PrintRune('\n')
}
