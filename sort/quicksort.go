package sort

import (
	"golang.org/x/exp/constraints"
)

// constraints.Ordered 대소비교가 가능한 타입만

func QuickSort[T constraints.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	// 왼쪽에는 pivot보다 작은값, 오른쪽에는 pivot보다 큰 값
	idx := patition(arr)
	QuickSort(arr[:idx])
	QuickSort(arr[idx+1:])
}

func patition[T constraints.Ordered](arr []T) int {
	if len(arr) <= 1 {
		return 0
	}
	pivot := arr[0]
	i := 1
	j := len(arr) - 1
	for {
		for i < len(arr) && arr[i] < pivot {
			i++
		}
		for j > 0 && arr[j] > pivot {
			j--
		}
		if i >= j {
			arr[0], arr[i-1] = arr[i-1], arr[0]
			return i - 1 // i와 j 가 같거나 역전될때까지 반복하다가, 작은값과 큰값이 양옆으로 분리되었을 때의 그 값을 피봇으로 설정
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func IsSorted[T constraints.Ordered](arr []T) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
