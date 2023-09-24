package linkedlist

// Node 链表节点
type Node struct {
	Data interface{}
	Next *Node
}

// LinkList 单向链表
type LinkList struct {
	Head   *Node
	Length int
}

// Init 初始化链表
func (l *LinkList) Init() {
	l.Head = &Node{}
	l.Length = 0
}

// IsEmpty 判断链表是否为空
func (l *LinkList) IsEmpty() bool {
	return l.Length == 0
}

// Insert 插入节点
func (l *LinkList) Insert(data interface{}) {
	node := &Node{Data: data}
	node.Next = l.Head.Next
	l.Head.Next = node
	l.Length++
}

// InitWithArray 使用数组初始化链表
func (l *LinkList) InitWithArray(arr []interface{}) {
	l.Init()
	for _, v := range arr {
		l.Insert(v)
	}
}

// Delete 删除第n个节点
func (l *LinkList) Delete(n int) {
	if n > l.Length {
		return
	}
	node := l.Head
	for i := 0; i < n-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	l.Length--
}

// DeleteWithData 删除第一个值为data的节点
func (l *LinkList) DeleteWithData(data interface{}) {
	dummy := &Node{Next: l.Head}
	node := dummy
	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			l.Length--
			break
		}
		node = node.Next
	}
}

// DeleteAllWithData 删除所有值为data的节点
func (l *LinkList) DeleteAllWithData(data interface{}) {
	dummy := &Node{Next: l.Head}
	node := dummy
	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			l.Length--
		} else {
			node = node.Next
		}
	}
}

// Get 获取第n个节点的值
func (l *LinkList) Get(n int) interface{} {
	if n > l.Length {
		return nil
	}
	node := l.Head.Next
	for i := 0; i < n-1; i++ {
		node = node.Next
	}
	return node.Data
}

// GetIndex 获取第一个值为data的节点的序号
func (l *LinkList) GetIndex(data interface{}) int {
	node := l.Head.Next
	for i := 0; i < l.Length; i++ {
		if node.Data == data {
			return i + 1
		}
		node = node.Next
	}
	return -1
}

// GetNode 获取第一个值为data的节点
func (l *LinkList) GetNode(data interface{}) *Node {
	node := l.Head.Next
	for i := 0; i < l.Length; i++ {
		if node.Data == data {
			return node
		}
		node = node.Next
	}
	return nil
}

// LastNode 获取链表的最后一个节点
func (l *LinkList) LastNode() *Node {
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	return node
}

// Update 更新第n个节点的值
func (l *LinkList) Update(n int, data interface{}) {
	if n > l.Length {
		return
	}
	node := l.Head.Next
	for i := 0; i < n-1; i++ {
		node = node.Next
	}
	node.Data = data
}
