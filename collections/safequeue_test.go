package collections

import (
	"runtime"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSafeQueue(t *testing.T) {
	t.Parallel()

	q := NewSafeQueue([]int{1, 2, 3, 4, 5})
	arr := []int{}

	copy := q.Array()

	initCount := q.Len()

	count := 0
	for !q.IsEmpty() {
		x, _ := q.Dequeue()
		arr = append(arr, x)
		count++
	}

	assert.Equal(t, count, len(arr))
	assert.Equal(t, initCount, len(arr))
	assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5})
	assert.ElementsMatch(t, copy, []int{1, 2, 3, 4, 5})
}

func TestNilSafeQueue(t *testing.T) {
	t.Parallel()

	var q *SafeQueue[int]

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

	i = q.Len()
	assert.Zero(t, i)
	assert.Nil(t, q)

	q = q.Enqueue(1)

	i, b = q.Peek()
	assert.Equal(t, i, 1)
	assert.True(t, b)
	assert.NotNil(t, q)
	assert.Equal(t, 1, q.Len())

	q = NewSafeQueue[int](nil)

	i, b = q.Dequeue()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.NotNil(t, q)

	i, b = q.Peek()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.NotNil(t, q)
}

func TestSafeQueueThreadSafety(t *testing.T) {
	t.Parallel()

	const total = 1000

	ss := NewSafeQueue[int](nil)

	wg := &sync.WaitGroup{}

	cpus := runtime.NumCPU()
	wg.Add(cpus)
	for i := 0; i < cpus; i++ {
		go func() {
			for j := 0; j < total; j++ {
				ss.Enqueue(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	arr := ss.Array()
	sort.Ints(arr)

	assert.Equal(t, total*cpus, ss.Len())
}
