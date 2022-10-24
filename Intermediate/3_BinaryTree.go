package main

import "fmt"

func main() {

	root := &BTNode{1,&BTNode{2,&BTNode{4, nil, nil}, nil},&BTNode{3,&BTNode{5, nil, nil},&BTNode{6, nil, nil}}}

	root.dfs(0)

}

type BTNode struct {
	val int
	left *BTNode
	right *BTNode
}

func (node *BTNode) dfs(space int) {
	if node == nil {
		return;
	}
	for i := (2 * space); i > 0; i-- {
		fmt.Print("-")
	}
	fmt.Println(node.val)
	node.left.dfs(space + 1)
	node.right.dfs(space + 1)
}