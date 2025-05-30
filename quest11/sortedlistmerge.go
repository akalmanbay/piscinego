package piscine

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	lfe := l1
	for lfe.Next != nil {
		lfe = lfe.Next
	}
	lfe.Next = l2
	return ListSort(l1)
}
