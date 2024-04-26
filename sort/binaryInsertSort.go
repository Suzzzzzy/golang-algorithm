package sort

import (
	"golang.org/x/exp/constraints"
	"log"
)

// 이진 - 둘로 나눠 한 쪽만 탐색하는 것
// 이미 정렬된 배열에 값을 삽입할 때 유용한 정렬
// 값이 계속 추가되는 자료구조에 유용

// 배열에 새로 값을 추가할때, 기존 배열의 가운데 값과 비교 -> 왼쪽 Or 오른쪽 삽입 확인

func BinaryInsertSort[T constraints.Ordered](sorted []T, val T) []T {
	idx := findInsert(sorted, val)
	return append(sorted[:idx], append([]T{val}, sorted[idx:]...)...)
}

func findInsert[T constraints.Ordered](sorted []T, val T) int {
	if len(sorted) == 0 {
		return 0 // 중간값보다 계속 작아질때 결국 맨 처음에 삽입
	}
	mid := len(sorted) / 2
	log.Println(mid)
	if sorted[mid] < val {
		return findInsert(sorted[mid+1:], val) + mid + 1 // 중간값으로 나눈 배열의 중간값 으로 나눈 배열의 중간값... 이 한개가 남으면 mid 다음 index
	} else {
		return findInsert(sorted[:mid], val)
	}
}
