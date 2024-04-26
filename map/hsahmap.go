package _map

import (
	"hash/crc32"
)

const arraySize = 3571

type hashData[T any] struct {
	key   string
	value T
}

type HashMap[T any] struct {
	arr [arraySize][]hashData[T]
}

func hashfn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//type HashMap[T any] struct {
//	arr [arraySize]hashData[T]
//}
//func (h *HashMap[T]) AddSimple(key string, value T) {
//	hash := hashfn(key)
//	h.arr[hash%arraySize] = hashData[T]{
//		key:   key,
//		value: value,
//	}
//}
//func (h *HashMap[T]) GetSimple(key string) (T, bool) {
//	hash := hashfn(key)
//	val := h.arr[hash%arraySize]
//	return val.value, key == val.key
//}

// Add 키를 해쉬한 것을 array에 저장하는 것
func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	var hd = hashData[T]{
		key:   key,
		value: value,
	}
	// slice로 바꿨기 때문에 추가하는 방식으로 해야함
	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hd) // 원래있는 데이터 뒤에 새로운 값 추가 - 충돌나더라도 데이터를 뒤에 다 추가할 수 있도록
}

// Get 키를 해쉬한 것을 가져와서 그 값이 실제 Key와 같은지 확인하는 것
func (h *HashMap[T]) Get(key string) (T, bool) { // bool은 값이 있는지 없는지 여부
	hash := hashfn(key)
	for _, hd := range h.arr[hash%arraySize] {
		if hd.key == key {
			return hd.value, true
		}
	}
	var tempT T
	return tempT, false
}

/* arr [arraySize]hashData[T]
Hash 충돌이 일어날 수 있다
다른 입력을 넣었는데 같은 hash가 나온다면 같은 위치에 데이터가 저장되기 때문

*/

/*
Hash Map 속도
추가 O(1)
조회 O(1)
삭제 O(1)
-> 항상 일정하다는 말
왜냐면 key를 hash하는 hash function의 속도이기 때문
요소의 개수와 상관 없다 요소의 개수가 백만개가 되어도 key hash function 함수의 속도로 일정

Sparse한 자료구조: 듬성듬성 저장되기 때문
<-> Dense한 자료구조: 빽빽하게 저장된다 => 캐쉬 친화적 구조이다 => 캐쉬 조회하는데 성능이 좋다
*/
