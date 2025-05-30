package piscine

func ReverseMenuIndex(menu []string) []string {
	size := len(menu)
	res := make([]string, size)
	j := len(menu) - 1
	for i := range menu {
		res[i] = menu[j]
		j--
	}

	return res
}
