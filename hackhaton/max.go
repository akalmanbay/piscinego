package piscine

func Max(a []int) int {
	n := 0
	for _, val := range a {
		if n < val {
			n = val
		}
	}
	return n
}
