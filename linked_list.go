package main

// LinkedList represents a double linked list.
type LinkedList struct {
	head *LinkedListElement
}

// LinkedListElement represents a single element of a LinkedList type.
type LinkedListElement struct {
	key  int
	next *LinkedListElement
	prev *LinkedListElement
}

// Insert appends an element in the list and makes it
// the new head of the linked list. The current head becomes
// the next element on the list.
//
// `O(1)` complexity.
func (l *LinkedList) Insert(key int) {
	e := LinkedListElement{key: key, next: l.head}

	// If there's an actual head, update it to point
	// to the previous element.
	if l.head != nil {
		l.head.prev = &e
	}

	l.head = &e
}

// Search finds the first ocurrence of the key within a list.
// Returns `nil` if not found.
//
// `O(n)` complexity.
func (l *LinkedList) Search(key int) *LinkedListElement {
	e := l.head

	for e != nil {
		if e.key == key {
			return e
		}

		e = e.next
	}

	return nil
}

// Delete removes a particular element from the linked list.
//
// `O(1)` complexity.
func (l *LinkedList) Delete(e *LinkedListElement) {
	if e.prev != nil {
		e.prev.next = e.next // If it's nil, is propagated
	} else {
		// If previous is null, then this is the head, thus
		// we must modify the linked list.
		l.head = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev // If it's nil, is propagated
	}
}
