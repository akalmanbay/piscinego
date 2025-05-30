package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printRune(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		data := make([]byte, 8)

		_, err := os.Stdin.Read(data)
		if err == nil {
			printRune(string(data))
		}
	}

	for i := 1; i < len(args); i++ {
		f, err := os.ReadFile(args[i])
		if err != nil {
			printRune("ERROR: " + err.Error() + "\n")
			os.Exit(1)

		}
		printRune(string(f))
	}
}
