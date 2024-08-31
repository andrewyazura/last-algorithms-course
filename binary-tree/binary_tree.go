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

func CompareTrees[T comparable](a *Node[T], b *Node[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return CompareTrees(a.left, b.left) && CompareTrees(a.right, b.right)
}

func Walk[T comparable](node *Node[T], ch chan *T) {
	if node == nil {
		ch <- nil
		return
	}

	ch <- &node.value
	Walk(node.left, ch)
	Walk(node.right, ch)
}

func CompareTreesAsync[T comparable](a *Node[T], b *Node[T]) bool {
	ch_a := make(chan *T)
	ch_b := make(chan *T)

	go func() {
		Walk(a, ch_a)
		close(ch_a)
	}()

	go func() {
		Walk(b, ch_b)
		close(ch_b)
	}()

	for i := range ch_a {
		j := <-ch_b

		if i == nil && j == nil {
			continue
		}

		if i == nil || j == nil {
			return false
		}

		if *i != *j {
			return false
		}
	}

	return true
}
