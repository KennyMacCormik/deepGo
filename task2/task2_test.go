package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CircularQueue struct {
	values          []int
	begin, end, num int
}

func (q *CircularQueue) next(i *int) {
	if *i+1 >= cap(q.values) {
		*i = 0
	} else {
		*i++
	}
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{make([]int, size), 0, 0, 0}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Empty() {
		q.values[q.begin] = value
		q.num++
		return true
	}
	if q.num+1 <= cap(q.values) {
		q.next(&q.end)
		q.num++
		q.values[q.end] = value
		return true
	}
	return false
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.next(&q.begin)
	q.num--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.begin]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.end]
}

func (q *CircularQueue) Empty() bool {
	return q.num == 0
}

func (q *CircularQueue) Full() bool {
	return q.num == cap(q.values)
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

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
