package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicQueue(t *testing.T) {
	t.Parallel()

	q := Queue[int]([]int{1, 2, 3, 4, 5})
	arr := []int{}

	initCount := len(q)

	count := 0
	for !q.IsEmpty() {
		x, _ := q.Dequeue()
		arr = append(arr, x)
		count++
	}

	assert.Equal(t, count, len(arr))
	assert.Equal(t, initCount, len(arr))
	assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5})
}

func TestNilQueue(t *testing.T) {
	t.Parallel()

	var q Queue[int]

	assert.True(t, q.IsEmpty())
	assert.Nil(t, q)

	i, b := q.Dequeue()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, q)

	i, b = q.Peek()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, q)

	q.Enqueue(1)

	i, b = q.Peek()
	assert.Equal(t, i, 1)
	assert.True(t, b)
	assert.NotNil(t, q)
}
