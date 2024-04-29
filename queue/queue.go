package queue

import (
	"algorithm/double_linked_list"
)

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

// 큐 두개로 스택 구현
type StackUsingQueue struct {
	mainQueue []int
	subQueue  []int
}

func (s *StackUsingQueue) Push(item int) {
	s.mainQueue = append(s.mainQueue, item)
}

func (s *StackUsingQueue) Pop() int {
	if len(s.mainQueue) == 0 {
		return -1 // stack에 아무것도 없으면 -1 반환
	}

	for len(s.mainQueue) > 1 {
		s.subQueue = append(s.subQueue, s.mainQueue[0])
		s.mainQueue = s.mainQueue[1:]
	}

	poped := s.mainQueue[0]
	s.mainQueue = []int{}

	s.mainQueue, s.subQueue = s.subQueue, s.mainQueue // 서로 값 바꾸기
	return poped
}

// 스택 두개로 큐 구현
type QueueUsingStack struct {
	mainStack []int
	subStack  []int
}

func (q *QueueUsingStack) EnQueue(item int) {
	q.mainStack = append(q.mainStack, item)
}

func (q *QueueUsingStack) DeQueue() int {
	if len(q.subStack) == 0 {
		for len(q.mainStack) > 0 {
			lastIndex := len(q.mainStack) - 1
			q.subStack = append(q.subStack, q.mainStack[lastIndex])
			q.mainStack = q.mainStack[:lastIndex]
		}
	}

	if len(q.subStack) > 0 {
		lastIndex := len(q.subStack) - 1
		dequeued := q.subStack[lastIndex]
		q.subStack = q.subStack[:lastIndex]
		return dequeued
	}
	return -1
}
