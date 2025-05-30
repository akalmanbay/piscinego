package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	res := 0
	tens := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 48 && s[i] <= 57 {
			digit := int((s[i] - '0'))
			res = res + digit*tens
			tens = tens * 10
		} else {
			return -1
		}
	}
	if res >= 1 && res <= 26 {
		return res
	} else {
		return -1
	}
}

func IsUpper(args []string) bool {
	if args[1] == "--upper" {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	isUpper := IsUpper(args)

	if isUpper {
		args = args[2:]
	} else {
		args = args[1:]
	}

	for i := 0; i < len(args); i++ {
		value := args[i]
		res := Atoi(value)
		if isUpper && res >= 0 {
			z01.PrintRune(rune(res + 64))
		} else if res >= 0 {
			z01.PrintRune(rune(res + 96))
		} else {
			z01.PrintRune(32)
		}

	}
	z01.PrintRune('\n')
}
