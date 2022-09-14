package collections

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSafeStack(t *testing.T) {
	t.Parallel()

	s := NewSafeStack(1, 2, 3, 4, 5)
	arr := []int{}

	initCount := s.Len()

	count := 0
	for !s.IsEmpty() {
		x, _ := s.Pop()
		arr = append(arr, x)
		count++
	}

	assert.Equal(t, count, len(arr))
	assert.Equal(t, initCount, len(arr))
	assert.ElementsMatch(t, arr, []int{5, 4, 3, 2, 1})
}

func TestNilSafeStack(t *testing.T) {
	t.Parallel()

	var s *SafeStack[int]

	assert.True(t, s.IsEmpty())
	assert.Nil(t, s)

	i, b := s.Pop()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, s)

	i, b = s.Peek()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, s)

	i = s.Len()
	assert.Zero(t, i)
	assert.Nil(t, s)

	arr := s.Slice()
	assert.Empty(t, arr)
	assert.NotNil(t, arr)
	assert.Nil(t, s)

	s = s.Push(1)

	i, b = s.Peek()
	assert.Equal(t, i, 1)
	assert.True(t, b)
	assert.NotNil(t, s)
	assert.Equal(t, 1, s.Len())
	assert.ElementsMatch(t, s.Slice(), []int{1})

	s = NewSafeStack[int]()

	i, b = s.Pop()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.NotNil(t, s)

	i, b = s.Peek()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.NotNil(t, s)
}

func TestSafeStackThreadSafety(t *testing.T) {
	t.Parallel()

	const total = 1000

	ss := NewSafeStack[int]()

	wg := &sync.WaitGroup{}

	cpus := runtime.NumCPU()
	wg.Add(cpus)
	for i := 0; i < cpus; i++ {
		go func() {
			for j := 0; j < total; j++ {
				ss.Push(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, total*cpus, ss.Len())
}
