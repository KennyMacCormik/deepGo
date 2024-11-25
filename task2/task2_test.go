package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	left, right int
	// need to implement
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{make([]int, 0, size), 0, 0}
}

func (q *CircularQueue) Push(value int) bool {
	if q.right-q.left == cap(q.values) {
		return false
	}

	if q.left != 0 {
		q.values[q.left-1] = value
		q.left--
		return true
	}

	q.values = append(q.values, value)
	q.right++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.values == nil || q.left == q.right {
		return false
	}

	q.left++
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.left]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.right-1]
}

func (q *CircularQueue) Empty() bool {
	return q.left == q.right
}

func (q *CircularQueue) Full() bool {
	return len(q.values) != 0 && q.right == cap(q.values) && q.left == 0
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 4, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
