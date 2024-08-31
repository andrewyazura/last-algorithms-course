package linked_lists

import (
	"errors"
)

type DoubleNode[T comparable] struct {
	value T
	next  *DoubleNode[T]
	prev  *DoubleNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head   *DoubleNode[T]
	tail   *DoubleNode[T]
	length int
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	new_node := &DoubleNode[T]{value: value}
	l.length++

	if l.head == nil {
		l.head = new_node
		l.tail = new_node
		return
	}

	new_node.next = l.head
	l.head.prev = new_node
	l.head = new_node
}

func (l *DoublyLinkedList[T]) InsertAt(value T, index int) error {
	if index > l.length || index < 0 {
		return errors.New("index out of range")
	} else if index == l.length {
		l.Append(value)
		return nil
	} else if index == 0 {
		l.Prepend(value)
		return nil
	}

	new_node := &DoubleNode[T]{value: value}
	l.length++

	current := l.getNode(index)

	new_node.next = current
	new_node.prev = current.prev

	current.prev.next = new_node
	current.next.prev = new_node

	return nil
}

func (l *DoublyLinkedList[T]) Append(value T) {
	new_node := &DoubleNode[T]{value: value}
	l.length++

	if l.tail == nil {
		l.head = new_node
		l.tail = new_node
		return
	}

	new_node.prev = l.tail
	l.tail.next = new_node
	l.tail = new_node
}

func (l *DoublyLinkedList[T]) Remove(value T) error {
	current := l.head
	for current != nil {
		if current.value == value {
			break
		}
		current = current.next
	}

	if current == nil {
		return errors.New("value not in list")
	}

	l.removeNode(current)
	return nil
}

func (l *DoublyLinkedList[T]) Get(index int) *T {
	switch node := l.getNode(index); node {
	case nil:
		return nil
	default:
		return &node.value
	}
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) *T {
	node := l.getNode(index)
	if node == nil {
		return nil
	}

	l.removeNode(node)
	return &node.value
}

func (l *DoublyLinkedList[T]) getNode(index int) *DoubleNode[T] {
	if index >= l.length || index < 0 {
		return nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

func (l *DoublyLinkedList[T]) removeNode(node *DoubleNode[T]) {
	l.length--

	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {
		l.head = node.next
	}

	if node == l.tail {
		l.tail = node.prev
	}

	node.next = nil
	node.prev = nil
}
