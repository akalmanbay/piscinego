package piscine

func JumpOver(str string) string {
	hell := ""
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			hell += string(str[i])
		}
	}
	hell += "\n"
	return hell
}
