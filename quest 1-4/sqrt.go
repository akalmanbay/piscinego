package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	} else {
		for i := 2; i < nb; i++ {
			if nb/i == i && nb%i == 0 {
				return i
			}
		}
	}
	return 0
}
