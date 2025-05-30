package piscine

func Unmatch(a []int) int {
	for _, v := range a {
		matches := 0
		for _, v2 := range a {
			if v == v2 {
				matches++
			}
		}

		if matches%2 != 0 {
			return v
		}
	}
	return -1
}
