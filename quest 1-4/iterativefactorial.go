package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb >= 1 {
		return nb * IterativeFactorial(nb-1)
	} else {
		return 0
	}
}
