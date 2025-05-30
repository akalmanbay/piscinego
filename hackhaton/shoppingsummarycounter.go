package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	res := map[string]int{}
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			res[word]++
			word = ""
		} else {
			word = word + string(str[i])
		}
	}
	res[word]++
	return res
}
