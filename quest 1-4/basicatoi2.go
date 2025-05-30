package piscine

func BasicAtoi2(s string) int {
	res := 0
	cnt := len(s) - 1
	for i, val := range s {
		temp := int(val - '0')
		if !(temp >= 0 && temp <= 9) {
			return 0
		}
		tens := 1
		for x := 0; x < cnt-i; x++ {
			tens = tens * 10
		}
		res = res + temp*tens

	}
	return res
}
