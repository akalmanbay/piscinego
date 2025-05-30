package piscine

func IsLower(s string) bool {
	for _, value := range s {
		if value <= 'z' && value >= 'a' {
			continue
		} else {
			return false
		}
	}
	return true
}
