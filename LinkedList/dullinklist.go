package linkedlist

type DulNode struct {
	Data interface{}
	Next *DulNode
	Pre  *DulNode
}

type DulLinkList struct {
	Head   *DulNode
	Length int
}

func (l *DulLinkList) Init() {
	l.Head = &DulNode{}
	l.Length = 0
}

func (l *DulLinkList) GetLength() int {
	return l.Length
}

func (l *DulLinkList) IsEmpty() bool {
	return l.Length == 0
}
