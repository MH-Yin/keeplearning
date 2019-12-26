package structure

type LRUList struct {
	l   *List
	cap int
}

// NewLruList returns List which cap
func NewLruList(cap int) *LRUList {
	if cap == 0 {
		return nil
	}
	return &LRUList{
		l:   NewList(),
		cap: cap,
	}
}

func (lr *LRUList) Insert(value interface{}) {
	if lr.l.Size == lr.cap {
		lr.removeOldest()
	}
	lr.l.Insert(value)
}

func (lr *LRUList) Remove(n *Node) {
	lr.l.Remove(n)
}

func (lr *LRUList) GetSize() int {
	return lr.l.Size
}

func (lr *LRUList) GetFirstNode() *Node {
	return lr.l.FirstNode
}

func (lr *LRUList) GetLastNode() *Node {
	return lr.l.LastNode
}

func (lr *LRUList) removeOldest() {
	removeNode := lr.l.FirstNode
	lr.l.Remove(removeNode)
}
