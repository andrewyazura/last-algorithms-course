package binary_tree

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func TraversePreOrder[T comparable](node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	result := make([]T, 0)
	result = append(result, node.value)
	result = append(result, TraversePreOrder(node.left)...)
	result = append(result, TraversePreOrder(node.right)...)

	return result
}

func TraverseInOrder[T comparable](node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	result := make([]T, 0)
	result = append(result, TraverseInOrder(node.left)...)
	result = append(result, node.value)
	result = append(result, TraverseInOrder(node.right)...)

	return result
}

func TraversePostOrder[T comparable](node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	result := make([]T, 0)
	result = append(result, TraversePostOrder(node.left)...)
	result = append(result, TraversePostOrder(node.right)...)
	result = append(result, node.value)

	return result
}

func TraverseBreadthFirst[T comparable](node *Node[T]) []T {
	result := make([]T, 0)
	queue := []*Node[T]{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.value)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return result
}
