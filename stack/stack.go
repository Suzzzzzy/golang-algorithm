package stack

import "algorithm/double_linked_list"

type Stack[T any] struct {
	l *double_linked_list.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &double_linked_list.LinkedList[T]{},
	}
}

func (s *Stack[T]) Push(val T) {
	s.l.PushBack(val)
}

func (s *Stack[T]) Pop() T {
	return s.l.PopBack().Value
}
