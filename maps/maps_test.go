package maps_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModifyPtr(t *testing.T) {
	m := make(map[string]*int)
	a := 11
	m["Calin"] = &a

	b := m["Calin"]
	*b = 13

	fmt.Printf("%#v", m)

	require.Equal(t, 13, *m["Calin"])
}

func TestAccess(t *testing.T) {
	m := make(map[string]int)
	require.Equal(t, 0, m["not_existing"])

	val, ok := m["not_existing"]
	require.Equal(t, 0, val)
	require.Equal(t, false, ok)

	m["existing"] = 0
	val, ok = m["existing"]
	require.Equal(t, 0, val)
	require.Equal(t, true, ok)

	m2 := make(map[int]int, 6)
	require.Equal(t, 5, len(m2))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, val := range nums {
		complement := target - val
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[val] = i
	}
	return nil
}
