package main

import (
	"errors"
)

// Queue represents a queue data structure for integers.
type Queue struct {
	data []int
	head int
	tail int
}

// Enqueue adds an element to a queue with `O(1)` complexity.
func (queue *Queue) Enqueue(value int) error {
	// When head = tail + 1 we're under an overflow. However, when the head
	// is at 0 and tail at the end of the queue, is an edge case that also
	// falls under the overflow case.
	if queue.head == queue.tail+1 || (queue.head == 0 && queue.tail == len(queue.data)-1) {
		return errors.New("overflow")
	}

	queue.data[queue.tail] = value

	if queue.tail == len(queue.data)-1 {
		queue.tail = 0
	} else {
		queue.tail++
	}

	return nil
}

// Dequeue removes a value (integer) from the queue and
// returns it.
func (queue *Queue) Dequeue() (int, error) {
	if queue.head == queue.tail {
		return 0, errors.New("underflow")
	}

	value := queue.data[queue.head]

	if queue.head == len(queue.data)-1 {
		queue.head = 0
	} else {
		queue.head++
	}

	return value, nil
}
