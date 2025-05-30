package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	idx := 0
	i := l
	for i != nil {

		if idx == pos {
			return i
		}
		idx++
		i = i.Next
	}
	return nil
}
