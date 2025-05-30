package piscine

func ListReverse(l *List) {
	if l.Tail == nil && l.Head == nil {
		l = nil
		return
	}

	l.Tail = l.Head

	var prev *NodeL
	prev = nil

	i := l.Head

	for i != nil {
		next := i.Next
		i.Next = prev

		prev = i
		i = next

	}
	l.Head = prev
}
