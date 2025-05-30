package piscine

func ActiveBits(n int) int {
	cnt := 0
	for ; n > 0; n = n / 2 {
		if n%2 == 1 {
			cnt++
		}
	}
	return cnt
}
