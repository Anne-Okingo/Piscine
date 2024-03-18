package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	char := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = char
		l.Tail = char
	} else {
		l.Tail.Next = char
		l.Tail = char
	}
}
