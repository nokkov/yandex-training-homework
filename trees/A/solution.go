package A

// package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// // iterate to find right place => insert
// func addNode(root *Node, data int) {
// 	if root.data == data {
// 		return
// 	}

// 	if root.data > data {
// 		if root.left == nil {
// 			root.left = &Node{data, nil, nil}
// 			return
// 		} else {
// 			addNode(root.left, data)
// 		}
// 	} else {
// 		if root.right == nil {
// 			root.right = &Node{data, nil, nil}
// 			return
// 		} else {
// 			addNode(root.right, data)
// 		}
// 	}
// }

// func maxDepth(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return max(maxDepth(root.left), maxDepth(root.right)) + 1
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	} else {
// 		return y
// 	}
// }

// func main() {

// 	var num int

// 	fmt.Scan(&num)
// 	root := &Node{num, nil, nil}

// 	for {
// 		fmt.Scan(&num)

// 		if num == 0 {
// 			break
// 		}

// 		addNode(root, num)
// 	}

// 	fmt.Print(maxDepth(root))
// }
