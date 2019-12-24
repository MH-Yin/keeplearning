// list && lruList
package structure

type List interface {
	Insert(value interface{})
	Remove(node Node)
	GetSize() int
	GetFirstNode() Node
	GetLastNode() Node
}

type Node interface {
	GetValue() interface{}
	GetPreNode() Node
	GetNextNode() Node
	IsNil() bool
}

type list struct {
	firstNode *node
	lastNode  *node
	size      int
}

type node struct {
	nextNode *node
	preNode  *node
	value    interface{}
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) GetPreNode() Node {
	return n.preNode
}

func (n *node) GetNextNode() Node {
	return n.nextNode
}

func (n *node) IsNil() bool {
	return n == nil
}

func NewList() List {
	return new(list)
}

func (l *list) Insert(value interface{}) {
	newNode := &node{
		value: value,
	}
	if l.size != 0 {
		newNode.preNode = l.lastNode
		l.lastNode.nextNode = newNode
	} else {
		l.firstNode = newNode
	}
	l.lastNode = newNode
	l.size++
}

func (l *list) GetFirstNode() Node {
	return l.firstNode
}

func (l *list) GetLastNode() Node {
	return l.lastNode
}

func (l *list) Remove(n Node) {
	no, ok := n.(*node)
	if !ok {
		return
	}
	defer func() {
		l.size--
	}()
	// Remove first node
	if no.preNode == nil {
		l.firstNode = no.nextNode
		if l.firstNode != nil {
			no.nextNode.preNode = nil
		}
		return
	}

	// Remove last node
	if no.nextNode == nil {
		l.lastNode = no.preNode
		if l.lastNode != nil {
			no.preNode.nextNode = nil
		}
		return
	}

	// node is middle in list
	no.preNode.nextNode = no.nextNode
	no.nextNode.preNode = no.preNode
}

func (l *list) GetSize() int {
	return l.size
}

type lruList struct {
	l   *list
	cap int
}

func NewLruList(cap int) List {
	if cap == 0 {
		return nil
	}
	return &lruList{
		l:   &list{},
		cap: cap,
	}
}

func (lr *lruList) Insert(value interface{}) {
	if lr.l.size == lr.cap {
		lr.removeOldest()
	}
	lr.l.Insert(value)
}

func (lr *lruList) Remove(n Node) {
	lr.l.Remove(n)
}

func (lr *lruList) GetSize() int {
	return lr.l.size
}

func (lr *lruList) GetFirstNode() Node {
	return lr.l.firstNode
}

func (lr *lruList) GetLastNode() Node {
	return lr.l.lastNode
}

func (lr *lruList) removeOldest() {
	removeNode := lr.l.firstNode
	lr.l.Remove(removeNode)
}
