package linklist

// Node 链表的节点
type Node struct {
	next  *Node
	value interface{}
}

// NewNode 新建一个链表节点
func NewNode(v interface{}) *Node {
	return &Node{next: nil, value: v}
}

func (n *Node) GetNode() *Node {
	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

// SingleLinkList 单链表的结构
type SingleLinkList struct {
	head *Node
	len  uint
}

// NewSingleLinkList 建立一个链表
func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{&Node{nil, nil}, 0}
}

// InsertAfter 向指定的节点后面插入数据
func (s *SingleLinkList) InsertAfter(n *Node, v interface{}) bool {
	if n == nil {
		return false
	}
	// 新的节点
	newNode := NewNode(v)
	oldNode := n.next
	n.next = newNode
	newNode.next = oldNode
	s.len++
	return true
}

//InsertBefore 向指定的节点前面插入数据
func (s *SingleLinkList) InsertBefore(n *Node, v interface{}) bool {
	if n == nil || s.head == n {
		return false
	}
	// 判断是否是head
	cur := s.head.next
	head := s.head
	for cur != nil {
		if cur == n {
			break
		}
		head = cur
		cur = cur.next
	}
	newNode := NewNode(v)
	head.next = newNode
	newNode.next = cur
	s.len++
	return true
}
