package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		isUpper := value <= 'Z' && value >= 'A'
		isLower := value <= 'z' && value >= 'a'
		isNumerical := value <= '9' && value >= '0'
		if isUpper || isLower || isNumerical {
			continue
		} else {
			return false
		}
	}
	return true
}
