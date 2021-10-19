package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (result []int) {
	var post func(node *Node)
	post = func(node *Node) {
		if root == nil {
			return
		}
		for _, c := range node.Children {
			post(c)
		}
		result = append(result, node.Val)
	}
	post(root)
	return
}

func main() {
}
