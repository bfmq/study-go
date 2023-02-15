package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	isExist := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if n, ok := isExist[node]; ok {
			return n
		}

		cloneNode := &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		isExist[node] = cloneNode
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(node)
}
