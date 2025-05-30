package piscine

func TrimAtoi(s string) int {
	tens := 1
	sign := 1
	isFirstNumber := false

	for _, value := range s {
		if value >= 48 && value <= 57 {
			tens = tens * 10
		}
	}
	res := 0
	for _, value := range s {
		if value >= 48 && value <= 57 {
			tens = tens / 10
			res = res + int(value-'0')*tens
			isFirstNumber = true

		} else if value == 45 && !isFirstNumber {
			sign = -1
		}
	}

	return res * sign
}
