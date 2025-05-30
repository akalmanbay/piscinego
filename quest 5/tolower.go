package piscine

func ToLower(s string) string {
	var res []rune
	// s1 := []rune(s)

	for _, value := range s {
		if value <= 90 && value >= 65 {
			res = append(res, value+32)
		} else {
			res = append(res, value)
		}
	}
	return string(res)
}
