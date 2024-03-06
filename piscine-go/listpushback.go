package piscine

func ListPushBack(l *List, data interface{}) {
	// n := &List{Head: &NodeL{Data: data}}
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil && l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
