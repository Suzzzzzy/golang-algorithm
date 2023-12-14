package binarytree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value any
	Left  *TreeNode
	Right *TreeNode
}
