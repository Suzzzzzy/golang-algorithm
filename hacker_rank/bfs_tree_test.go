package hacker_rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
문제: https://www.hackerrank.com/challenges/bfsshortreach/problem
[1,2] [1,3] 이런식으로 노드가 입력되고 최소의 시작점으로 부터 각 노드의 거리를 출력
*/

type node struct {
	value  int32
	childs []*node
}

func addEdge(nodeList map[int32]*node, u, v int32) {
	nodeList[u].childs = append(nodeList[u].childs, nodeList[v])
	nodeList[v].childs = append(nodeList[v].childs, nodeList[u])
}

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	nodeList := make(map[int32]*node) // 꼭 만들어주기

	for i := int32(1); i <= n; i++ {
		nodeList[i] = &node{value: i}
	}

	for _, edge := range edges {
		addEdge(nodeList, edge[0], edge[1])
	}

	distances := make([]int32, n+1)
	for i := int32(1); i <= n; i++ {
		distances[i] = -1
	}
	distances[s] = 0

	queue := []*node{nodeList[s]}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range node.childs {
			if distances[child.value] == -1 {
				distances[child.value] = distances[node.value] + 6
				queue = append(queue, child)
			}
		}

	}
	result := []int32{}
	for i := int32(1); i <= n; i++ {
		if i != s {
			result = append(result, distances[i])
		}
	}

	return result
}

func TestBfsTree(t *testing.T) {
	n := int32(4)
	m := int32(2)
	edges := [][]int32{
		{1, 2},
		{1, 3},
	}
	result := bfs(n, m, edges, 1)
	expectedResult := []int32{6, 6, -1}
	println(result)
	assert.Equal(t, result, expectedResult)
}

