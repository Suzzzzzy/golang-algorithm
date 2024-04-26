package _map

import "golang.org/x/exp/constraints"

type SparseSet[TKey constraints.Ordered, TValue any] struct {
	dense  []Element[TKey, TValue]
	sparse map[TKey]int // key와 array에 저장된 index 저장
}

// NewSparseSet new로 하는 이유는 map이 레퍼런서 타입이다 보니 초기화를 해주어야함 - 그렇지 않으면 nullpoint error 발생
func NewSparseSet[TKey constraints.Ordered, TValue any]() *SparseSet[TKey, TValue] {
	return &SparseSet[TKey, TValue]{
		sparse: make(map[TKey]int),
	}
}

func (s *SparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].Value = value // 이미 key에 해당하는 값이 있다면 덮어쓰기
		return
	}

	s.dense = append(s.dense, Element[TKey, TValue]{
		Key:   key,
		Value: value,
	})
	s.sparse[key] = len(s.dense) - 1 // dense에 추가했기 때문에 dense의 마지막 인덱스 값
}

func (s *SparseSet[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	if idx, ok := s.sparse[key]; ok {
		value = s.dense[idx].Value
		found = true
	} else {
		found = false
	}
	return value, found
}

func (s *SparseSet[TKey, TValue]) Remove(key TKey) bool {
	if idx, ok := s.sparse[key]; ok {
		last := len(s.dense) - 1
		if idx < last { // 내가지우려는 요소가 맨 마지막이면 바꿔줄 필요없이 삭제만 하면 되기 때문에 맨 마지막이 아닌 경우 로직
			s.dense[idx] = s.dense[last]
			s.sparse[s.dense[last].Key] = idx
		}
		s.dense = s.dense[:last]
		delete(s.sparse, key)
		return true
	}
	return false
}

type Iterator[TKey constraints.Ordered, TValue any] struct {
	dense []Element[TKey, TValue]
	idx   int
}

func (s *SparseSet[TKey, TValue]) Iterator() *Iterator[TKey, TValue] {
	return &Iterator[TKey, TValue]{
		dense: s.dense, // slice 카피 이므로 length, capacity, 시작 포인터만 카피되므로 O(1)
		idx:   0,
	}
}

func (i *Iterator[TKey, TValue]) IsEnd() bool {
	return i.idx >= len(i.dense)
}

func (i *Iterator[TKey, TValue]) Next() {
	i.idx++
}

func (i *Iterator[TKey, TValue]) Get() Element[TKey, TValue] {
	return i.dense[i.idx]
}
