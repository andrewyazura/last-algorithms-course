package linked_lists

import "errors"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type List[T comparable] struct {
	head   *Node[T]
	length int
}

func (l *List[T]) Append(value T) {
	new_node := &Node[T]{value: value}
	l.length++

	if l.head == nil {
		l.head = new_node
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = new_node
}

func (l *List[T]) Prepend(value T) {
	new_node := &Node[T]{value: value}
	new_node.next = l.head
	l.head = new_node
	l.length++
}

func (l *List[T]) InsertAt(value T, index int) (*Node[T], error) {
	if index > l.length {
		return nil, errors.New("index out of range")
	}

	new_node := &Node[T]{value: value}
	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	new_node.next = current.next
	current.next = new_node
	l.length++

	return new_node, nil
}

func (l *List[T]) Remove(value T) error {
	if l.head == nil {
		return errors.New("list is empty")
	}

	if l.head.value == value {
		e := l.head
		l.head = l.head.next

		e.next = nil // avoid memory leaks
		l.length--
		return nil
	}

	current := l.head
	for current.next != nil {
		if current.next.value == value {
			e := current.next
			current.next = e.next

			e.next = nil // avoid memory leaks
			l.length--
			return nil
		}

		current = current.next
	}

	return errors.New("not found")
}

func (l *List[T]) RemoveAt(value T) error {
	return nil
}
