package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	for i := l; i.Next != nil; i = i.Next {
		for j := l; j.Next != nil; j = j.Next {
			if j.Data > j.Next.Data {
				j.Data, j.Next.Data = j.Next.Data, j.Data
			}
		}
	}
	return l
}
