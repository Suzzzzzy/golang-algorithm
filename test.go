package golang_algorithm

import (
	"sort"
	"strings"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i := 0
	j := 0
	idx := 0
	result := make([]int, len(left)+len(right))
	for i < len(left) || j < len(right) {
		var leftMerge bool
		if i >= len(left) {
			leftMerge = false
		} else if j >= len(right) {
			leftMerge = true
		} else {
			leftMerge = left[i] < right[j]
		}

		if leftMerge {
			result[idx] = left[i]
			i++
		} else {
			result[idx] = right[j]
			j++
		}
		idx++
	}
	return result
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	idx := pivot(arr)
	quickSort(arr[:idx])
	quickSort(arr[idx+1:])
}

func pivot(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	i := 1
	j := len(arr) - 1

	for {
		for i < len(arr) && arr[i] < arr[0] {
			i++
		}
		for j > 0 && arr[j] > arr[0] {
			j--
		}
		if i >= j {
			arr[0], arr[i-1] = arr[i-1], arr[0]
			return i - 1
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type element struct {
	key   string
	value int
}

type sortedMap struct {
	arr []element
}

func (s *sortedMap) add(key string, value int) {
	idx := sort.Search(len(s.arr), func(i int) bool {
		return s.arr[i].key >= key
	})

	if idx < len(s.arr) && s.arr[idx].key == key {
		s.arr[idx].value = value
		return
	}

	s.arr = append(s.arr[:idx], append([]element{{key, value}}, s.arr[idx:]...)...)
}

func (s *sortedMap) get(key string) (value int, ok bool) {
	idx := sort.Search(len(s.arr), func(i int) bool {
		return s.arr[i].key >= key
	})
	if idx < len(s.arr) && s.arr[idx].key == key {
		return s.arr[idx].value, true
	}
	var defaultValue int
	return defaultValue, false
}

func (s *sortedMap) remove(key string) bool {
	idx := sort.Search(len(s.arr), func(i int) bool {
		return s.arr[i].key >= key
	})
	if idx < len(s.arr) && s.arr[idx].key == key {
		s.arr = append(s.arr[:idx], s.arr[idx+1:]...)
		return true
	}
	return false
}

// report 당근 문제
func report(input []string) string {
	reportCounts := make(map[string]int)

	for _, reported := range input {
		if _, ok := reportCounts[reported]; !ok {
			reportCounts[reported]++
		}
	}

	blackList := ""
	maxCount := 0
	for reported, count := range reportCounts {
		if count > maxCount {
			maxCount = count
			parts := strings.Fields(reported)
			blackList = parts[1]
		}
	}
	return blackList
}

type treeNode struct {
	value  string
	childs []*treeNode
}

func (t *treeNode) add(val string) *treeNode {
	n := &treeNode{
		value: val,
	}
	t.childs = append(t.childs, n)
	return n
}

func (t *treeNode) bfs(fn func(val string)) {
	queue := make([]*treeNode, 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		fn(front.value)

		for _, n := range front.childs {
			queue = append(queue, n)
		}
	}
}

func (t *treeNode) dfs(fn func(val string)) {
	stack := make([]*treeNode, 0)
	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.value)

		stack = append(stack, last.childs...)
	}
}

type graph struct {
	list    map[int][]int
	visited map[int]bool
	order   []int
}

func (g *graph) add(from, to int) {
	g.list[from] = append(g.list[from], to)
	g.list[to] = append(g.list[to], from)
}

func (g *graph) listSort() {
	for _, v := range g.list {
		sort.Ints(v)
	}
}

func (g *graph) dfs(v int) {
	g.visited[v] = true
	g.order = append(g.order, v)

	for _, i := range g.list[v] {
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

func (g *graph) bfs(v int) {
	queue := []int{v}
	g.visited[v] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		g.order = append(g.order, current)

		for _, neighbor := range g.list[current] {
			if !g.visited[neighbor] {
				g.visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
