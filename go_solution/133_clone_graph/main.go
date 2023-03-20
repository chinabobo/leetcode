package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

// 递归克隆，传入已经访问过的元素作为过滤条件
// DFS

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	// 已经访问过直接返回
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}
