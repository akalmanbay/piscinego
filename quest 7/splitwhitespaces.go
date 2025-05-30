package piscine

func SplitWhiteSpaces(s string) []string {
	var ans []string
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
		} else {
			if str != "" {
				ans = append(ans, str)
			}
			str = ""
		}
	}
	if str != "" {
		ans = append(ans, str)
	}
	return ans
}
