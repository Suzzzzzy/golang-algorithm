package tree

type TreeNode[T any] struct {
	Value  T
	Childs []*TreeNode[T] // 여러개의 자식 노드를 가질 수 있음
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}
	t.Childs = append(t.Childs, n)
	return n
}

func (t *TreeNode[T]) Preorder(fn func(val T)) {
	if t == nil {
		return
	}
	fn(t.Value) // 본인
	// 자식들 순회
	for _, n := range t.Childs {
		n.Preorder(fn)
	}
}

func (t *TreeNode[T]) Postorder(fn func(val T)) {
	if t == nil {
		return
	}
	// 자식들 순회
	for _, n := range t.Childs {
		n.Preorder(fn)
	}
	fn(t.Value) // 본인
}

// BFS 너비 우선 탐색
func (t *TreeNode[T]) BFS(fn func(val T)) {
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t) //root를 먼저 큐에 집어 넣는다

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:] // 큐에서 맨 첫번째 것만 빼오는 것

		fn(front.Value)

		for _, n := range front.Childs {
			queue = append(queue, n)
		}

	}
}

// DFS 깊이 우선 탬색(재귀함수 대신 스택 사용)
func (t *TreeNode[T]) DFS(fn func(val T)) {
	stack := []*TreeNode[T]{}
	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.Value)

		stack = append(stack, last.Childs...)
		stack = append(last.Childs, stack...)
	}
}
