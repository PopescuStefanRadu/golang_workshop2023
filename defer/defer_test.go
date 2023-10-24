package defer_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func deferExampleWithClosure() {
	name := "John"
	defer func() {
		fmt.Println(name)
	}()

	name = "Georgina"
	defer func() {
		fmt.Println(name)
	}()
}

func deferExampleWithVariableCapture() {
	name := "John"
	defer func(str string) {
		fmt.Println(str)
	}(name)

	name = "Georgina"
	defer func(str string) {
		fmt.Println(str)
	}(name)
}

func TestDefer(t *testing.T) {
	deferExampleWithClosure()
	// Georgina
	// Georgina

	fmt.Println()

	deferExampleWithVariableCapture()
	// Georgina
	// John
}

func plusOne() (x int) {
	defer func() {
		x++
	}()
	return 1
}

func plusOneBis() (x int) {
	defer func() {
		x++
	}()
	return 1
}

func TestPlusOne(t *testing.T) {
	require.Equal(t, plusOne(), 1)
	require.Equal(t, plusOneBis(), 1)
}
