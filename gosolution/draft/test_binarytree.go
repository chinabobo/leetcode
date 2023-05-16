package main

import (
	"fmt"
	"github.com/chinabobo/leetcode/util/binarytree"
)

func main() {
	root := binarytree.BuildTree([]int{1, 2, 3, 4, 0, 5, 6})
	fmt.Println(binarytree.PreorderTraversalNonRecursive(root))

	// result := binarytree.ConvertToArr(root)
	// fmt.Println(result)
	// fmt.Println(BFS(root))
	// fmt.Println(BFS2(root))
}
