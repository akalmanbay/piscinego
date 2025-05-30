package piscine

func ToUpper(s string) string {
	var res []rune
	// s1 := []rune(s)

	for _, value := range s {
		if value <= 122 && value >= 97 {
			res = append(res, value-32)
		} else {
			res = append(res, value)
		}
	}
	return string(res)
}
