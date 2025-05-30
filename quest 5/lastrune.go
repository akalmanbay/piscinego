package piscine

func LastRune(s string) rune {
	res := []rune(s)
	return res[len(s)-1]
}
