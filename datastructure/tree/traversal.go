package tree

import "fmt"

func RecursivePreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	RecursivePreOrderTraversal(root.left)
	RecursivePreOrderTraversal(root.right)
}

func NonRecursivePreOrderTraversal(root *Node) {
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			fmt.Printf("%v ", root.val)
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			root = root.right
		}
	}
}

func RecursiveInOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	RecursiveInOrderTraversal(root.left)
	fmt.Printf("%v ", root.val)
	RecursiveInOrderTraversal(root.right)
}

func NonRecursiveInOrderTraversal(root *Node) {
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			fmt.Printf("%v ", root.val)
			stack = stack[0 : len(stack)-1]
			root = root.right
		}
	}
}

func RecursivePostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	RecursivePostOrderTraversal(root.left)
	RecursivePostOrderTraversal(root.right)
	fmt.Printf("%v ", root.val)
}

func NonRecursivePostOrderTraversal(root *Node) {
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		if root != nil { // 递归
			stack = append(stack, root)
			root = root.left
		} else { // 回溯
			if stack[len(stack)-1].right != nil { // 有右子树，继续递归
				root = stack[len(stack)-1].right
			} else { // 没有右子树，开始打印
				for len(stack) > 1 && stack[len(stack)-1] == stack[len(stack)-2].right { // 打印并向上回溯，直到到达还没有进行递归的地方
					fmt.Printf("%v ", stack[len(stack)-1].val)
					stack = stack[0 : len(stack)-1]
				}
				fmt.Printf("%v ", stack[len(stack)-1].val)
				stack = stack[0 : len(stack)-1]
				if len(stack) > 0 { // 继续递归
					root = stack[len(stack)-1].right
				}
			}
		}
	}
}

func NonRecursivePostOrderTraversalSimple(root *Node) {
	stack := make([]*Node, 0)
	printStack := make([]int, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			printStack = append(printStack, root.val)
			stack = append(stack, root)
			root = root.right
		} else {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			root = root.left
		}
	}
	for i := len(printStack) - 1; i >= 0; i-- {
		fmt.Printf("%v ", printStack[i])
	}
}

func HierarchicalTraversal(root *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", temp.val)
		if temp.left != nil {
			queue = append(queue, temp.left)
		}
		if temp.right != nil {
			queue = append(queue, temp.right)
		}
	}
}
