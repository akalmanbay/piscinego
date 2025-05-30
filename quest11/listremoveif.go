package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	var prev *NodeL
	prev = nil
	i := l.Head

	for i != nil {
		if i.Data == data_ref {
			if prev == nil {
				l.Head = i.Next
			} else {
				prev.Next = i.Next
			}
		} else {
			prev = i
		}
		i = i.Next
	}
}
