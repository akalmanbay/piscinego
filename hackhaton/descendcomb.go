package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for m := '9'; m >= '0'; m-- {
				for n := '9'; n >= '0'; n-- {
					if ((i == m && j > n) || (i > m)) && !(i == m && j == n) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(m)
						z01.PrintRune(n)
						if !(i == '0' && j == '1' && m == '0' && n == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
