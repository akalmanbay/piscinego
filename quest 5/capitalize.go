package piscine

func isLetter(value byte) bool {
	return (value >= 48 && value <= 57) || (value >= 65 && value <= 90) || (value >= 97 && value <= 122)
}

func isUpper(value byte) bool {
	return value >= 65 && value <= 90
}

func isLower(value byte) bool {
	return value >= 97 && value <= 122
}

func toLower(s byte) byte {
	return s + 32
}

func toUpper(s byte) byte {
	return s - 32
}

func Capitalize(s string) string {
	res := ""
	prev := false
	current := false
	for i := 0; i < len(s); i++ {
		current = isLetter(s[i])

		if prev && current && isUpper(s[i]) {
			res = res + string(toLower(s[i]))
		} else if !prev && current && isLower(s[i]) {
			res = res + string(toUpper(s[i]))
		} else {
			res = res + string(s[i])
		}

		prev = current

	}
	return res
}
