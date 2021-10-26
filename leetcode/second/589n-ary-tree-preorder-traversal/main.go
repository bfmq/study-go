package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (result []int) {
	// 定义辅助递归函数
	var pre func(node *Node)
	pre = func(node *Node) {
		if node == nil {
			// 节点为空则停止递归
			return
		}
		// 前序遍历先加自己
		result = append(result, node.Val)
		// 循环子节点一个个递归
		for _, n := range node.Children {
			pre(n)
		}
	}
	// root节点入递归
	pre(root)
	return
}

func main() {
}
