package smallproblem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(ints []int) (res []int) {
	m := make(map[int]struct{}) // empty struct holds no memory
	for _, v := range ints {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = struct{}{}
		}
	}
	return
}

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, removeDuplicates([]int{1, 2, 3, 3, 4, 4}))
}
