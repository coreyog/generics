package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicStack(t *testing.T) {
	t.Parallel()

	st := Stack[int]([]int{1, 2, 3, 4, 5})
	arr := []int{}

	initCount := len(st)

	count := 0
	for !st.IsEmpty() {
		x, _ := st.Pop()
		arr = append(arr, x)
		count++
	}

	assert.Equal(t, count, len(arr))
	assert.Equal(t, initCount, len(arr))
	assert.ElementsMatch(t, arr, []int{5, 4, 3, 2, 1})
}

func TestNilStack(t *testing.T) {
	t.Parallel()

	var st Stack[int]

	assert.True(t, st.IsEmpty())
	assert.Nil(t, st)

	i, b := st.Pop()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, st)

	i, b = st.Peek()
	assert.Zero(t, i)
	assert.False(t, b)
	assert.Nil(t, st)

	st.Push(1)

	i, b = st.Peek()
	assert.Equal(t, i, 1)
	assert.True(t, b)
	assert.NotNil(t, st)
}
