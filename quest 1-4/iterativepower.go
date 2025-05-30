package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 || nb == 0 {
		return 0
	} else if power == 1 {
		return nb
	} else {
		return nb * IterativePower(nb, power-1)
	}
}
