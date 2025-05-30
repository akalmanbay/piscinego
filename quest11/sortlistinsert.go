package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return nil
	}
	newNode := &NodeI{Data: data_ref}
	lfe := l
	for lfe.Next != nil {
		lfe = lfe.Next
	}
	lfe.Next = newNode
	ListSort(l)
	return l
}
