package piscine

func Join(elems []string, sep string) string {
	res := ""
	for i, value := range elems {
		res = res + string(value)
		if i < len(elems)-1 {
			res = res + sep
		}
	}
	return res
}
