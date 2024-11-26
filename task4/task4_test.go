package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	key, val *int
	left     *OrderedMap
	right    *OrderedMap
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{} // need to implement
}

func (m *OrderedMap) Insert(key, value int) {
	f := func(item **OrderedMap) {
		if *item == nil {
			k, v := key, value
			*item = &OrderedMap{key: &k, val: &v}
			return
		}
		(*item).Insert(key, value)
	}
	// empty tree
	if m.key == nil && m.val == nil && m.left == nil && m.right == nil {
		k, v := key, value
		m.key, m.val = &k, &v
		return
	}
	// self
	if *m.key == key {
		*m.val = value
		return
	}
	// left
	if key < *m.key {
		f(&m.left)
	} else {
		f(&m.right) // right
	}
}

func (m *OrderedMap) Erase(key int) {
	var f func(item **OrderedMap)
	f = func(item **OrderedMap) {
		if *item == nil {
			return
		}
		if *(*item).key == key {
			// none
			if (*item).right == nil && (*item).left == nil {
				*item = nil
				return
			}
			// only left
			if (*item).right == nil && (*item).left != nil {
				*item = (*item).left
				return
			}
			// only right
			if (*item).right != nil && (*item).left == nil {
				*item = (*item).right
				return
			}
			// both
			rep := (*item).right
			repParent := *item
			for rep.left != nil {
				repParent = rep
				rep = rep.left
			}
			(*item).key = rep.key
			(*item).val = rep.val
			// Delete actual rep
			if rep.right != nil {
				if repParent.left == rep {
					repParent.left = rep.right
				} else {
					repParent.right = rep.right
				}
			} else {
				if repParent.left == rep {
					repParent.left = nil
				} else {
					repParent.right = nil
				}
			}
		}
		if key < *(*item).key {
			f(&(*item).left)
		} else {
			f(&(*item).right)
		}
	}

	f(&m)
}

func (m *OrderedMap) Contains(key int) bool {
	if m == nil || m.key == nil {
		return false
	}
	// self
	if *m.key == key {
		return true
	}
	// left
	if key < *m.key {
		return m.left != nil && m.left.Contains(key)
	}
	// right
	return m.right != nil && m.right.Contains(key)
}

func (m *OrderedMap) Size() int {
	if m.key == nil && m.val == nil && m.left == nil && m.right == nil {
		return 0
	}

	left, right := 0, 0

	if m.left != nil {
		left = m.left.Size()
	}
	if m.right != nil {
		right = m.right.Size()
	}

	return 1 + left + right
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m == nil || m.key == nil {
		return
	}
	//left
	if m.left != nil {
		m.left.ForEach(action)
	}
	//self
	action(*m.key, *m.val)
	//right
	if m.right != nil {
		m.right.ForEach(action)
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
