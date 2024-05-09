package backjun

// https://www.acmicpc.net/problem/10157

func ticketNum(C, R, K int) (int, int) {
	di := []int{0, +1, 0, -1} // 방향 순서에 따른 x좌표값의 변화
	dj := []int{+1, 0, -1, 0} // 방향 순서에 따른 y좌표값의 변화
	tmp := 1
	direction := 0
	x, y := 0, 0

	arr := make([][]int, C)
	for i := range arr {
		arr[i] = make([]int, R)
	}

	if R*C < K {
		print(0)
	}
	for tmp < C*R {
		if tmp == K {
			break
		} else {
			arr[x][y] = 1
			nextX := x + di[direction]
			nextY := y + dj[direction]
			if nextX < C && nextY < R { // 다음 점위 범위 내에 있고
				if arr[nextX][nextY] != 1 { // 방문한 점이 아니라면
					x = nextX // 이동하기
					y = nextY
					tmp += 1
				} else { // 방문했던 점이라면
					direction = (direction + 1) % 4 // 방향 바꾸기
				}
			} else { // 다음 점이 범위 밖이라면
				direction = (direction + 1) % 4 // 방향 바꾸기
			}
		}
	}
	return x, y
}
