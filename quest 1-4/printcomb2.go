package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var w rune = 48

	for ; w <= 57; w++ {
		var x rune = 48
		for ; x <= 57; x++ {
			var y rune
			var z rune
			if x < 57 {
				y = w
				z = x + 1

			} else {
				y = w + 1
				z = 48

			}
			for ; y <= 57; y++ {
				for ; z <= 57; z++ {
					z01.PrintRune(w)
					z01.PrintRune(x)
					z01.PrintRune(' ')
					z01.PrintRune(y)
					z01.PrintRune(z)
					if !(w == 57 && x == 56 && y == 57 && z == 57) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}

				}
				z = 48
			}
		}
	}
}
