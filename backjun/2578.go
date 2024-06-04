package main

import (
	"bufio"
	"fmt"
	"os"
)

// check: 빙고가 몇 개 인지 체크
func check(arr [][]int) int {
	cnt := 0
	// 세로 채크
	for i := 0; i < 5; i++ {
		column := 0
		for j := 0; j < 5; j++ {
			if arr[j][i] == 0 {
				column++
			}
		}
		if column == 5 {
			cnt++
		}
	}
	// 가로 체크
	for i := 0; i < 5; i++ {
		row := 0
		for j := 0; j < 5; j++ {
			if arr[i][j] == 0 {
				row++
			}
		}
		if row == 5 {
			cnt++
		}
	}
	// 대각선
	diaLeft := 0
	for i := 0; i < 5; i++ {
		if arr[i][i] == 0 {
			diaLeft++
		}
	}
	if diaLeft == 5 {
		cnt++
	}
	diaRight := 0
	for i := 0; i < 5; i++ {
		if arr[i][4-i] == 0 {
			diaRight++
		}
	}
	if diaRight == 5 {
		cnt++
	}

	return cnt
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	arr := make([][]int, 5) // make 로 만든것은 슬라이스 - 전달할때 포인터 전달 되므로 함수 내에서 요소 직접 변환 가능
	for i := range arr {
		arr[i] = make([]int, 5)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var n int
			fmt.Fscan(reader, &n)
			arr[i][j] = n
		}
	}
	var speak []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var n int
			fmt.Fscan(reader, &n)
			speak = append(speak, n)
		}
	}
	for i := 0; i < 25; i++ {
		s := speak[i]
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if arr[y][x] == s {
					arr[y][x] = 0 // 방문처리
					cnt := check(arr)
					if cnt >= 3 {
						fmt.Fprintln(writer, i+1)
						return
					}
				}
			}
		}

	}

}
