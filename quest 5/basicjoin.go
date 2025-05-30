package piscine

func BasicJoin(elems []string) string {
	res := ""
	for _, value := range elems {
		res = res + string(value)
	}
	return res
}
