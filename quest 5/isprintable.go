package piscine

func IsPrintable(s string) bool {
	for _, value := range s {
		if value > 126 || value < 32 {
			return false
		}
	}
	return true
}
