package collections

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSet(t *testing.T) {
	t.Parallel()

	s := ToSet(1, 2, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5)
	arr := []int{}

	initCount := len(s)

	count := 0
	for x := range s {
		arr = append(arr, x)
		count++
	}

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
	s.Delete(1) // tests that deleting a non-existent element doesn't panic
	assert.False(t, s.InSet(1))

	arr = s.Slice()
	sort.Ints(arr)
	assert.ElementsMatch(t, arr, []int{2, 3, 4, 5})
}

func TestStructSet(t *testing.T) {
	t.Parallel()

	type Person struct {
		Name string
		Age  int
	}

	alice := Person{
		Name: "Alice",
		Age:  30,
	}

	bob := Person{
		Name: "Bob",
		Age:  45,
	}

	corey := Person{
		Name: "Corey",
		Age:  33,
	}

	s := ToSet(alice, bob, corey, alice, bob, corey)

	danielle := Person{
		Name: alice.Name,
		Age:  alice.Age,
	}

	s.Add(danielle)

	assert.Equal(t, 3, len(s))
}

func TestNilSet(t *testing.T) {
	t.Parallel()

	var s Set[int]

	assert.True(t, s.IsEmpty())
	assert.Nil(t, s)

	s.Delete(1)
	assert.Nil(t, s)

	s.InSet(1)
	assert.Nil(t, s)

	s.Add(1)
	assert.NotNil(t, s)
	assert.True(t, s.InSet(1))
}

func TestStable(t *testing.T) {
	t.Parallel()

	arr := []int{1, 2, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	s := MakeStableSet(arr...)

	assert.NotNil(t, s)
	assert.Equal(t, 5, len(s))
	assert.ElementsMatch(t, s, []int{1, 2, 3, 4, 5})
}
