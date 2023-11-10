package buffer

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (s *SliceBuffer[T]) Write(data []T) {
	s.buf = append(s.buf, data...)
}

func (s *SliceBuffer[T]) Readable() int {
	return len(s.buf)
}
