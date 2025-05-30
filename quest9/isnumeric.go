package piscine

func IsNumeric(s string) bool {
	for _, value := range s {
		isNumerical := value <= '9' && value >= '0'
		if isNumerical {
			continue
		} else {
			return false
		}
	}
	return true
}
