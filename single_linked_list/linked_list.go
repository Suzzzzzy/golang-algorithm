package main

/*
Node는 연결 리스트에서 각각의 요소를 나타내는 기본 단위
연결 리스트는 데이터 요소를 저장하는 노드들이 서로 연결된 데이터 구조!
Node[T] 에서 T는 Generic 으로 정의되어 어떤 데이터 타입이든 저장할 수 있다
intNode := Node[int]{
	value: 20
	next: nil,
}
next는 다음 노드를 가리키는 포인터 -> 정수값 직접 할당할 수 없다
따라서 노드의 인스턴스를 만들고 그 노드의 포인터를 next에 할당해야 한다
*/

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil { // 아무것도 없는 상황에서는 Root가 Node
		l.root = node
		l.tail = node // 한개만 있는 상황이니 tail도 Node
		return
	}
	l.tail.next = node // 이미 Element가 들어있는 상황은 마지막 노드의 다음 노드를 새로만든 노드로 해주고
	l.tail = node      // 맨 마지막 꼬리가 바뀌었으니 tail을 노드로 바꾸어 준다
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root // 새로운 Node가 맨앞으로 추가되니까 노드의 다음이 기존 root
	l.root = node      // 그 다음 root를 새로운 노드로
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

// Count 노드는 사이즈가 없다. root 노드부터 하나씩 세어가는 것
// node의 개수만큼 돌게 되므로 O(N)
func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next { // if node != nil 이라는 조건 -> node가 nil이 아니면 다음 노드로 넘어가도록 무한 루프
		cnt++
	}
	return cnt
}

// Count2 두번째 방법 - push할 때마다 count 를 ++ 하는 방법도 있음
// O(1)
func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
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

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{ // 현재 노드의 next 노드를 newNode로 바꿔주어야 함
		Value: value,
	}
	originNext := node.next
	node.next = newNode
	newNode.next = originNext
	l.count++ // node가 추가되었기 때문
	// 버그 위험성: 리스트에 없는 노드인데 함수를 실행하는 경우 count 가 꼬여버림
}

// isIncluded node가 포함된 노드인지 확인
func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

// InsertBefore 이전 노드의 next를 새로운 값으로 바꾸고, 새로운 값의 next 노드를 현재 노드로 바꾸어 주면 된다
/* single linked list 로써 다음 노드만 알고 있지 이전 노드를 확인할 수는 없다
-> 이전 노드를 찾는 방법: 맨 앞에서부터 훑어보면서 다음 노드가 현재 노드인지 파악하여 알 수 있다(다음 노드가 현재 노드라면 해당 노드가 이전 노드)
*/
func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value) // 내가 추가하고싶은 노드가 root면 맨앞에 추가하는 pushFront와 같다
		return
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}
	prevNode.next = newNode
	newNode.next = node
	l.count++
}

// findPrevNode 이전 노드 찾기
func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

// PopFront 맨 앞의 것을 없애기: 두번째 노드가 첫번째 노드가 되면 된다
func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}
	l.root.next = nil // 참조 끊어주는 의미
	l.root = l.root.next
	l.count--
}
