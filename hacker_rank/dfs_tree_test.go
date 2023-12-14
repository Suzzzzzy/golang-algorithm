package hacker_rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 문제: https://www.hackerrank.com/challenges/ctci-connected-cell-in-a-grid/problem
// DFS: 깊이 우선 탐색 - 경로 탐색

func isValid(i, j, rows, cols int32, visited [][]bool, grid [][]int32) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols && !visited[i][j] && grid[i][j] == 1
}

// dfs 인자 한개를 입력하면 그 주변 상하좌우의 것을 모두 탐색하고 또 계속해서 탐색하는 함수
func dfs(i, j, rows, cols int32, visited [][]bool, grid [][]int32) int32 {
	direction := [][]int32{{-1, -1}, {-1, 0}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
	size := int32(1)
	visited[i][j] = true

	for _, dir := range direction {
		ni, nj := i+dir[0], j+dir[1]
		if isValid(ni, nj, rows, cols, visited, grid) {
			size += dfs(ni, nj, rows, cols, visited, grid)
		}
	}
	return size
}

func maxRegion(grid [][]int32) int32 {
	rows := int32(len(grid))
	if rows == 0 {
		return 0
	}
	cols := int32(len(grid[0]))

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	maxRegionSize := int32(0)

	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			regionSize := dfs(i, j, rows, cols, visited, grid)
			if regionSize > maxRegionSize {
				maxRegionSize = regionSize
			}
		}
	}
	return maxRegionSize
}

func TestDfsTree(t *testing.T) {
	grid := [][]int32{
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{1, 0, 0, 0},
	}
	expectedResult := 5
	result := maxRegion(grid)

	assert.Equal(t, expectedResult, result)
}
