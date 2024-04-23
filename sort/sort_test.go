package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	QuickSort(arr)
	assert.True(t, IsSorted(arr))
}

func TestMergeSort(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	sorted := MergeSort(arr)
	assert.True(t, IsSorted(sorted))
}

func TestBinaryInsertSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(100))
	}
	assert.True(t, IsSorted(arr), arr)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}
	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}
	MergeSort(arr)
}
