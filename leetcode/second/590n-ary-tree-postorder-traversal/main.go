package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (result []int) {
	// 定义辅助递归函数
	var post func(node *Node)
	post = func(node *Node) {
		if node == nil {
			// 节点为空则停止递归
			return
		}
		// 循环子节点一个个递归
		for _, n := range node.Children {
			post(n)
		}
		// 后序遍历最后加自己
		result = append(result, node.Val)
	}
	// root节点入递归
	post(root)
	return
}

func main() {
}
