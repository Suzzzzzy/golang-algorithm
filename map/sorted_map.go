package _map

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type Element[TKey constraints.Ordered, TValue any] struct {
	Key   TKey
	Value TValue
}

type SortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []Element[TKey, TValue]
}

func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})
	/*
		sort.Search : func 만족하는 최소한의 i(index) 값을 찾아줌
		작은 값부터 찾기 때문에
	*/
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr[idx].Value = value
		return
	}
	s.Arr = append(s.Arr[:idx],
		append([]Element[TKey, TValue]{
			{Key: key, Value: value},
		}, s.Arr[idx:]...)...)
	// 앞의 것에다가, (넣으려는 요소 + 뒤의 것) 합친것을 합치는 것
}

func (s *SortedMap[TKey, TValue]) Get(key TKey) (value TValue, ok bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key // key 값보다 크거나 같을 경우의 최소 인덱스
	})
	// key값에 맞는 값을 정확하게 찾았을 경우
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		return s.Arr[idx].Value, true
	}
	var defaultV TValue // 초기값
	return defaultV, false
}

func (s *SortedMap[TKey, TValue]) Remove(key TKey) (removed bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key // key 값보다 크거나 같을 경우의 최소 인덱스
	})
	// key값에 맞는 값을 정확하게 찾았을 경우
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...) // 해당 인덱스 부분만 빼고 왼쪽과 오른쪽을 합친다
		return true
	}
	return false
}