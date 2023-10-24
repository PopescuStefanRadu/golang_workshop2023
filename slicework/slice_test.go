package slicework_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[len(primes)-1:] /// last elem
	fmt.Println(s)

	var firstEl = primes[:1] /// first elem
	fmt.Println("firstEl", firstEl)

	last3Elems := primes[len(primes)-3:]
	fmt.Println("last3Elems", last3Elems)

	newSlice := []int{1}
	newSlice = append(newSlice, 2) // append
	fmt.Println("newSlice", newSlice)

	newSlice2 := []int{1}
	newSlice2 = append(newSlice2, 2, 3, 4)
	fmt.Println("newSlice2", newSlice2)

	appendAtStart := []int{1}
	appendAtStart = append([]int{0}, appendAtStart...)
	fmt.Println("appendAtStart", appendAtStart)

}

func TestSlicesFunctions(t *testing.T) {
	primes := []int{2, 3, 5, 7, 13, 11}
	slices.Sort(primes)
	fmt.Println(primes)

	slices.SortFunc(primes, func(a, b int) int {
		return a - b
	})

	//sort.Sort() //
}

func removeAtPos(elems []int, i int) []int {
	elems = append(elems[:i], elems[i+1:]...)
	return elems
}

func TestRemoveAtPos(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, removeAtPos([]int{1, 2, 3, 4}, 3))
	require.Equal(t, []int{1, 2, 4}, removeAtPos([]int{1, 2, 3, 4}, 2))
	require.Equal(t, []int{}, removeAtPos([]int{1}, 0))

	require.Panics(t, func() {
		removeAtPos([]int{}, 5)
	})
}

func TestSliceLiteral(t *testing.T) {
	sl := []int{1, 2, 3}
	require.Equal(t, 3, cap(sl))
	require.Equal(t, 3, len(sl))

	sl1 := []int{}
	require.Equal(t, 0, cap(sl1))
	require.Equal(t, 0, len(sl1))

	// Slice, pointer la un array, len, cap
	var sl2 []int
	require.Equal(t, 0, cap(sl2))
	require.Equal(t, 0, len(sl2))
}

func TestSliceGotchas(t *testing.T) {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	sl := names[:2] // len 2, cap 4
	require.Equal(t, 2, len(sl))
	require.Equal(t, 4, cap(sl))

	sl = append(sl, "test")
	require.Equal(t, [4]string{"John", "Paul", "test", "Ringo"}, names)
	require.Equal(t, []string{"John", "Paul", "test"}, sl)

	sl = append(sl, "test")
	sl = append(sl, "test")
	sl = append(sl, "test")
	sl = append(sl, "test")
	require.Equal(t, [4]string{"John", "Paul", "test", "test"}, names)
	require.Equal(t, []string{"John", "Paul", "test", "test", "test", "test", "test"}, sl)
}

func TestPerformance(_ *testing.T) {
	latestCapacity := 0
	sl := make([]int, 0, 0)
	for i := 0; i < 1<<20; i++ {
		sl = append(sl, 1)
		if latestCapacity != cap(sl) {
			latestCapacity = cap(sl)
			fmt.Println("latestCapacity", latestCapacity)
		}
	}
}

func TestMutabilityThroughAccessPatterns(t *testing.T) {
	type Vertex struct {
		X, Y int
	}

	vs := make([]Vertex, 3)
	for _, v := range vs {
		v.X++
	}
	assert.Equal(t, []Vertex{{X: 1}, {X: 1}, {X: 1}}, vs)

	for i := range vs {
		vs[i].X++
	}
	assert.Equal(t, []Vertex{{X: 1}, {X: 1}, {X: 1}}, vs)
}
