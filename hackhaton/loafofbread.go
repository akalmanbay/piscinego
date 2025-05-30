package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	}
	ans := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if count == 0 && ans != "" {
				ans += " "
			}
			ans += string(str[i])
			count++
			if count == 5 {
				count = 0
				i++
			}
		} else {
			continue
		}
	}
	return ans + "\n"
}
