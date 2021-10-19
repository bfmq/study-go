package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (result []int) {
	var pre func(node *Node)
	pre = func(node *Node) {
		if root == nil {
			return
		}
		result = append(result, node.Val)
		for _, c := range node.Children {
			pre(c)
		}
	}
	pre(root)
	return
}

func main() {
}
