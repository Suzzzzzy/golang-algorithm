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

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
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

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	n := l.root
	l.root = n.next
	if l.root != nil {
		l.root.prev = nil
	} else { // root가 nil이라는 것은 아무것도 없다는 것
		l.tail = nil
	}
	n.next = nil // 현재 노드(pop하려는 노드)의 next를 끊어준다
	l.count--
	return n
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	n := l.tail
	l.tail = n.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.root = nil
	}
	n.prev = nil
	l.count--
	return n

}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}
	n := &Node[T]{
		Value: val,
	}
	prevNode := node.prev
	node.prev = n

	n.prev = prevNode
	n.next = node

	if prevNode != nil {
		prevNode.next = n
	}
	if node == l.root {
		l.root = n
	}
	l.count++
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}
	n := &Node[T]{
		Value: val,
	}
	nextNode := node.next
	node.next = n

	n.next = nextNode // 새로운 노드의 다음 = 현재 노드의 Next
	n.prev = node     // 새로운 노드의 이전 = 현재 노드

	if nextNode != nil { // 새로운 노드의 다음이 nil이 아니라면
		nextNode.prev = n // 원래 Next 노드의 이전 노드 = 새로운 노드
	}
	if node == l.tail {
		l.tail = n // 마지막 노드인지 꼭 확인하는 이유 - noded의 tail 정보를 수정해야 하기 때문
	}

}

func (l *LinkedList[T]) Reverse() {
	newL := &LinkedList[T]{}
	if l.root == nil {
		return
	}
	for l.root != nil {
		n := l.PopFront()
		newL.PushFront(n.Value)
	}
	l.count = newL.count
	l.root = newL.root
	l.tail = newL.tail
}
