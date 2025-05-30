package piscine

func Rot14(s string) string {
	res := ""
	tmp := ""
	for _, v := range s {
		if (v >= 'A' && v <= 'L') || (v >= 'a' && v <= 'l') {
			tmp = string(v + 14)
		} else if (v > 'l' && v <= 'z') || (v > 'L' && v <= 'Z') {
			tmp = string(v - 12)
		} else {
			tmp = string(v)
		}

		res = res + (tmp)
	}
	return res
}
