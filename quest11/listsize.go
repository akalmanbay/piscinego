package piscine

func ListSize(l *List) int {
	cnt := 0
	i := l.Head
	for i != nil {
		i = i.Next
		cnt++
	}
	return cnt
}
