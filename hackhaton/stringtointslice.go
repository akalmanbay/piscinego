package piscine

func StringToIntSlice(str string) []int {
	hello := []int{}
	for _, val := range str {
		hello = append(hello, int(val))
	}
	if len(hello) == 0 {
		return nil
	}
	return hello
}
