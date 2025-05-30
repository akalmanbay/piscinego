package piscine

func Compact(ptr *[]string) int {
	res := []string{}

	cnt := 0
	for _, v := range *ptr {
		if v != "" {
			cnt++
			res = append(res, v)
		}
	}
	*ptr = res
	return cnt
}
