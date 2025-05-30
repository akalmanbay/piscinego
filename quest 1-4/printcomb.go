package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var x rune = 48
	for ; x <= 55; x++ {
		for z := x + 1; z <= 56; z++ {
			for y := z + 1; y <= 57; y++ {
				z01.PrintRune(x)
				z01.PrintRune(z)
				z01.PrintRune(y)
				if !(x == 55 && z == 56 && y == 57) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
