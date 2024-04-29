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

// Add 중간값을 보고 왼쪽인지 오른쪽인지 판단 - 또 그의 중간값을 보고 왼쪽인지 오른쪽인지 판단...
func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	/*
		sort.Search : func 만족하는 최소한의 i(index) 값을 찾아줌
		작은 값부터 찾기 때문에

	*/
	// array 가 아무것도 없는 초기 경우 체크 && 덮어씌우는 경우 체크
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
	if idx < len(s.Arr) && s.Arr[idx].Key == key { // 배열에 아무것도 없을 경우 idx =0 나옴 그래서 없는 배열에 array[0] 으로 접근하려니 에러발생
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

/*
Sorted Map 속도
1. 삽입: O(N)
index 뒤로 미뤄서 삽입하기 때문에 사실은 O(N-a) a는 앞의 요소 개수
그러나 상수는 무시가능하므로 O(N)
2. 조회: O(logN)
중간값 찾고, 그의 중간값 찾고.. 반씩 조회하므로 log2N
3. 삭제
요소 삭제하고 뒤의 것들을 땡겨주기 때문에

hash 맵이 훨씬 빠르지만 sorted map을 사용하는 이유:
- 정렬상태를 유지하기 위해
- Dense map 이므로 캐쉬 친화적
-
*/
