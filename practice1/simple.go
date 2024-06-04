package practice1

import (
	"sync"
	"time"
)

// 지수계산하는 재귀함수
func power(base int, exp int) int {
	// 지수가 0인 경우는 1
	if exp == 0 {
		return 1
	}

	return base * power(base, exp-1)
}

/* for 루프에서 반복될때마다 고루틴을 생성하여 게시물 조회를 병렬처리
 */

type Post struct {
	ID      int
	Content string
}

// mockFetchPost API로 부터 게시물을 가져오는 것을 시뮬레이션
func mockFetchPost(postID int) (Post, error) {
	time.Sleep(100 * time.Microsecond)
	return Post{ID: postID, Content: "content"}, nil
}

// fetchPosts 여러 게시물을 비동기적으로 가져온다
func fetchPosts(postIDs []int) ([]Post, error) {
	var wg sync.WaitGroup
	resultChan := make(chan Post, len(postIDs)) // Post 타입의 채널, 버퍼(크기)는 id 개수만큼
	errChan := make(chan error, len(postIDs))
	defer close(resultChan)
	defer close(errChan)

	for _, postID := range postIDs {
		wg.Add(1) // 완료해야하는 작업의 개수
		go func(id int) {
			defer wg.Done() // 작업이 완료되면 호출 -> 남은 작업의 개수를 줄여줌
			post, err := mockFetchPost(id)
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- post
		}(postID)
	}
	wg.Wait() // 전체 작업이 완료될때까지 대기

	var posts []Post
	var errs []error

	for len(resultChan) > 0 {
		select {
		case post := <-resultChan:
			posts = append(posts, post)
		case err := <-errChan:
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return posts, errs[0]
	}
	return posts, nil
}
