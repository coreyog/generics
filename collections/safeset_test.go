package collections

import (
	"runtime"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSafeSet(t *testing.T) {
	t.Parallel()

	s := NewSafeSet([]int{1, 2, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5})
	arr := []int{}

	initCount := s.Len()

	count := 0
	s.Range(func(x int) bool {
		arr = append(arr, x)
		count++

		return true
	})

	sort.Ints(arr)

	assert.Equal(t, count, len(arr))
	assert.Equal(t, initCount, len(arr))
	assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5})
	assert.False(t, s.InSet(0))
	assert.True(t, s.InSet(1))
	assert.True(t, s.InSet(2))
	assert.True(t, s.InSet(3))
	assert.True(t, s.InSet(4))
	assert.True(t, s.InSet(5))
	assert.False(t, s.InSet(6))
	assert.False(t, s.IsEmpty())

	s.Delete(1)
	assert.False(t, s.InSet(1))

	arr = s.Array()
	sort.Ints(arr)
	assert.ElementsMatch(t, arr, []int{2, 3, 4, 5})
}

func TestNilSafeSet(t *testing.T) {
	t.Parallel()

	var s *SafeSet[int]

	assert.True(t, s.IsEmpty())
	assert.Nil(t, s)

	s.Delete(1)
	assert.Nil(t, s)

	s.InSet(1)
	assert.Nil(t, s)

	assert.Zero(t, s.Len())
	assert.Nil(t, s)

	count := 0
	s.Range(func(x int) bool {
		count++
		return true
	})

	assert.Zero(t, count)
	assert.Nil(t, s)

	s = s.Add(1)
	assert.NotNil(t, s)
	assert.True(t, s.InSet(1))
}

func TestSafeSetThreadSafety(t *testing.T) {
	t.Parallel()

	const total = 1000

	ss := &SafeSet[int]{}

	wg := &sync.WaitGroup{}

	cpus := runtime.NumCPU()
	wg.Add(cpus)
	for i := 0; i < cpus; i++ {
		go func() {
			for j := 0; j < total; j++ {
				ss.Add(j)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	arr := ss.Array()
	sort.Ints(arr)

	assert.Equal(t, total, ss.Len())

	for i := 0; i < total; i++ {
		assert.True(t, ss.InSet(i))
		assert.Equal(t, i, arr[i])
	}
}
