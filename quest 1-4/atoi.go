package piscine

func Atoi(s string) int {
	res := 0
	cnt := len(s) - 1
	isNegative := 1
	for i, val := range s {
		temp := int(val - '0')
		if i == 0 && val == '-' {
			isNegative = -1
			continue
		} else if i == 0 && val == '+' {
			continue
		}
		if !(temp >= 0 && temp <= 9) {
			return 0
		}
		tens := 1
		for x := 0; x < cnt-i; x++ {
			tens = tens * 10
		}
		res = res + temp*tens

	}
	return res * isNegative
}
