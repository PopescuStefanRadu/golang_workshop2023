package pointers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Add(x int) {
	x++
}

func AddWithPtr(x *int) {
	if x != nil {
		*x++
	}
}

func TestAdd(t *testing.T) {
	x := 1
	Add(x)
	assert.Equal(t, x, 2)
	// False - pass by value
	// True

}

func TestAddWithPtr(t *testing.T) {
	x := 1
	ptrToX := &x
	AddWithPtr(ptrToX)
	assert.Equal(t, x, 2)
	assert.Equal(t, *ptrToX, 2)
}

func TestNPE(t *testing.T) {
	var s *int // s = nil
	require.Panics(t, func() {
		_ = *s // -> PANIC!!!
	})
	require.Panics(t, func() {
		*s = 10
	})
	// s = &i
	fmt.Println(*s)
}

//func Test
