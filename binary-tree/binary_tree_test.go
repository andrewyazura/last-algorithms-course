package binary_tree

import (
	"reflect"
	"testing"
)

func getTree() *Node[int] {
	r := Node[int]{value: 5}

	r.left = &Node[int]{value: 1}
	r.right = &Node[int]{value: 7}
	r.left.left = &Node[int]{value: 0}
	r.left.right = &Node[int]{value: 3}
	r.right.left = &Node[int]{value: 6}
	r.right.right = &Node[int]{value: 9}

	return &r
}

func TestPreOrder(t *testing.T) {
	r := getTree()
	path := TraversePreOrder(r)

	expected := []int{5, 1, 0, 3, 7, 6, 9}

	if !reflect.DeepEqual(path, expected) {
		t.Errorf("%v != %v", path, expected)
	}
}

func TestInOrder(t *testing.T) {
	r := getTree()
	path := TraverseInOrder(r)

	expected := []int{0, 1, 3, 5, 6, 7, 9}

	if !reflect.DeepEqual(path, expected) {
		t.Errorf("%v != %v", path, expected)
	}
}

func TestPostOrder(t *testing.T) {
	r := getTree()
	path := TraversePostOrder(r)

	expected := []int{0, 3, 1, 6, 9, 7, 5}

	if !reflect.DeepEqual(path, expected) {
		t.Errorf("%v != %v", path, expected)
	}
}

func TestBreadthFirst(t *testing.T) {
  r := getTree()
  path := TraverseBreadthFirst(r)

  expected := []int{5, 1, 7, 0, 3, 6, 9}

  if !reflect.DeepEqual(path, expected) {
    t.Errorf("%v != %v", path, expected)
  }
}
