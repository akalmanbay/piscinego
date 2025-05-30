package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var res []int
		return res
	}
	size := max - min
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = min
		min++
	}

	return res
}
