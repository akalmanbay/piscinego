package piscine

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0

	for _, s := range tab {
		if f(s) {
			cnt++
		}
	}
	return cnt
}
