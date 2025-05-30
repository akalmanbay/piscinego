package piscine

func IsUpper(s string) bool {
	for _, value := range s {
		if value <= 'Z' && value >= 'A' {
			continue
		} else {
			return false
		}
	}
	return true
}
