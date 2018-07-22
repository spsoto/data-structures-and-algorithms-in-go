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

// Swap positions of two elements in a linked list.
func (l *LinkedList) Swap(x *LinkedListElement, y *LinkedListElement) {
	if x.prev == nil {
		// It's the head. Make the swap.
		l.head = y
	} else {
		// The next of the previous is the new item.
		x.prev.next = y
	}

	if y.prev == nil {
		// It's the head. Make the swap.
		l.head = x
	} else {
		// The next of the previous is the new item.
		y.prev.next = x
	}

	if x.next != nil {
		// The previous of the next is the new item.
		x.next.prev = y
	}

	if y.next != nil {
		// The previous of the next is the new item.
		y.next.prev = x
	}

	// Now swap the pointers on the final item.
	x.prev, x.next, y.prev, y.next = y.prev, y.next, x.prev, x.next
}

// Reverse reverses the order of a linked list.
//
// `O(n)` complexity.
func (l *LinkedList) Reverse() {
	e := l.head

	for e != nil {
		if e.next == nil {
			l.head = e
		}

		e.next, e.prev, e = e.prev, e.next, e.next
	}
}
