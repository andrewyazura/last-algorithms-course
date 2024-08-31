package linked_lists

import (
	"testing"
)

func TestAppend(t *testing.T) {
  l := new(List[int])

  if l.head != nil {
    t.Errorf("l.head is not nil")
  }

  l.Append(10)

  if l.head == nil {
    t.Errorf("l.head is nil after appending")
  }

  if l.head.value != 10 {
    t.Errorf("l.head.value != 10")
  }

  l.Append(50)

  if l.head.next == nil {
    t.Errorf("l.head.next is nil")
  }

  if l.head.next.value != 50 {
    t.Errorf("l.head.next.value != 50")
  }
}

func TestPrepend(t *testing.T) {
  l := new(List[int])
  l.Prepend(50)

  if l.head == nil {
    t.Errorf("l.head is nil after appending")
  }

  if l.head.value != 50 {
    t.Errorf("l.head.value != 50")
  }

  l.Prepend(10)

  if l.head.value != 10 {
    t.Errorf("l.head.value != 10")
  }

  if l.head.next == nil {
    t.Errorf("l.head.next is nil")
  }

  if l.head.next.value != 50 {
    t.Errorf("l.head.next.value != 50")
  }
}

func TestInsertAt(t *testing.T) {
  l := new(List[int])
  l.Append(1)
  l.Append(2)
  l.Append(3)
  l.Append(4)
  l.Append(5)

  if l.head.next.next.value != 3 {
    t.Errorf("third node's value != 3")
  }

  l.InsertAt(6, 2)

  if l.head.next.next.value != 6 {
    t.Errorf("third node's value != 6")
  }
}

func TestRemove(t *testing.T) {
  l := new(List[int])
  l.Append(1)

  l.Remove(1)

  if l.head != nil {
    t.Errorf("l.head is not nil")
  }

  l.Append(1)
  l.Append(2)
  l.Append(3)

  l.Remove(2)

  if l.head.next.value != 3 {
    t.Errorf("l.head.next.value != 3")
  }
}
