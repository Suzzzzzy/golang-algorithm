package double_linked_list

// double linked list: 엣지가 양방향인 리스트

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n // root, tail 은 태그 같은 것 -> 꼬리표를 마지막 추가한 노드로 달아주는 것
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count() {
		return nil
	}
	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) PushFront(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}
	l.root.prev = n
	n.next = l.root
	l.root = n
	l.count++
	return
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	n := &Node[T]{
		Value: val,
	}
}
