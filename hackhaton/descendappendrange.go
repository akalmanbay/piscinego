package piscine

func DescendAppendRange(max, min int) []int {
	n := []int{}
	for i := max; i > min; i-- {
		n = append(n, i)
	}
	return n
}
