package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, v := range strs {
		res = res + v
		if i != len(strs)-1 {
			res = res + sep
		}

	}
	return res
}
