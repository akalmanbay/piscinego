package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = n + 9000000000000000000
		n = n * -1

	}
	if n < 0 {
		n = n * -1
		z01.PrintRune('-')
	}
	if n > 9 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}
