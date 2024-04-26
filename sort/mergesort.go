package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	i := 0   // left index
	j := 0   // right index
	idx := 0 // 정렬하여 합치는 배열 index
	rst := make([]T, len(left)+len(right))
	for i < len(left) || j < len(right) {
		var leftMerge bool
		if i >= len(left) {
			leftMerge = false // lefjt 쪽 merge 끝난경우
		} else if j >= len(right) {
			leftMerge = true // right 쪽 merge 끝난경우 = left 합쳐야함
		} else { // left, right 비교할 배열이 남아있고 진행중일때
			leftMerge = left[i] < right[j]
		}

		if leftMerge {
			rst[idx] = left[i]
			i++
		} else {
			rst[idx] = right[j]
			j++
		}
		idx++
	}
	return rst
}

func ArrayMergeSort[T constraints.Ordered](array1, array2 []T) []T {
	merged := make([]T, len(array1)+len(array2))
	i, j := 0, 0
	idx := 0

	for i < len(array1) || j < len(array2) {
		var leftMerge bool
		if i >= len(array1) {
			leftMerge = false
		} else if j >= len(array2) {
			leftMerge = true
		} else {
			leftMerge = array1[i] <= array2[j]
		}

		if leftMerge {
			merged[idx] = array1[i]
			i++
		} else {
			merged[idx] = array2[j]
			j++
		}
		idx++
	}
	return merged
}
