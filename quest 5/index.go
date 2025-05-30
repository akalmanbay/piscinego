package piscine

func Index(s string, toFind string) int {
	if len(toFind) > len(s) {
		return -1
	} else if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] && i+1 != len(s) {
			i_temp := i + 1
			j := 1
			for ; j < len(toFind); j++ {
				if s[i_temp] == toFind[j] {
					i_temp++
				} else {
					break
				}
			}
			if j == len(toFind) {
				return i
			}
		}
	}
	return -1
}
