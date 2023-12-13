package _map

import "hash/crc32"

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

// Add 키를 해쉬한 것을 array에 저장하는 것
func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	var hd = hashData[T]{
		key:   key,
		value: value,
	}
	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hd) // 충돌나더라도 데이터를 뒤에 다 추가할 수 있도록
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
