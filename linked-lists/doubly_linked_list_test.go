package linked_lists

import "testing"

func TestDoublyAppend(t *testing.T) {
  l := new(DoublyLinkedList[int])

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

  if l.head.next.prev != l.head {
    t.Errorf("l.head.next.prev != l.head")
  }
}

func TestDoublyPrepend(t *testing.T) {
  l := new(DoublyLinkedList[int])
  l.Prepend(50)

  if l.head == nil {
    t.Errorf("l.head is nil after prepending")
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

  if l.head.next.prev != l.head {
    t.Errorf("l.head.next.prev != l.head")
  }
}

func TestDoublyInsertAt(t *testing.T) {
  l := new(DoublyLinkedList[int])
  l.Append(10)
  l.Append(50)
  l.Append(100)

  l.InsertAt(25, 1)

  if l.head.next == nil {
    t.Errorf("l.head.next is nil")
  }

  if l.head.next.value != 25 {
    t.Errorf("l.head.next.value is %d, expected 25", l.head.next.value)
  }
}

func TestDoublyRemove(t *testing.T) {
  l := new(DoublyLinkedList[int])
  l.Append(10)
  l.Append(50)
  l.Append(100)

  l.Remove(50)

  if l.head.next == nil {
    t.Errorf("l.head.next is nil after removal")
  }

  if l.head.next.value != 100 {
    t.Errorf("l.head.next.value != 100")
  }

  if l.head.next.prev != l.head {
    t.Errorf("l.head.next.prev != l.head")
  }

  l.Remove(10)

  if l.head.value != 100 {
    t.Errorf("l.head.value != 100 after removing head")
  }

  if l.head.prev != nil {
    t.Errorf("l.head.prev is not nil after removing head")
  }

  l.Remove(100)

  if l.head != nil {
    t.Errorf("l.head is not nil after removing all elements")
  }
}
