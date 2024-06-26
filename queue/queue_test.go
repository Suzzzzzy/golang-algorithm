package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func TestPush2(t *testing.T) {
	s := NewSliceQueue[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func BenchmarkLinkedListQueue(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

// slice가 왠만하면 빠르다

func BenchmarkSliceQueue(b *testing.B) {
	s := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func TestStackUsingQueue(t *testing.T) {
	stack := StackUsingQueue{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	println("pop 요소 %v", stack.Pop())
	println("pop 요소 %v", stack.Pop())
	println("pop 요소 %v", stack.Pop())
	println(len(stack.mainQueue))
	println(len(stack.subQueue))
}

func TestQueueUsingStack(t *testing.T) {
	queue := QueueUsingStack{}
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	println("pop 요소", queue.DeQueue())
	println("pop 요소", queue.DeQueue())
	println("pop 요소", queue.DeQueue())
	println(len(queue.mainStack))
	println(len(queue.subStack))
}
