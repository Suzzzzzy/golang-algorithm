package queue

import "algorithm/double_linked_list"

type Queue[T any] struct {
	l *double_linked_list.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &double_linked_list.LinkedList[T]{},
	}
}

func (q *Queue[T]) Push(val T) {
	q.l.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	return q.l.PopFront().Value
}
