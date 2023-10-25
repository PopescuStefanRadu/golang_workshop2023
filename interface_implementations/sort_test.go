package interface_implementations_test

import (
	"esol/interface_implementations"
	"fmt"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func sortUsingSortSort(in []int) {
	// sortam intai dupa paritate, impar < par
	// pentru egalitate in paritate sortam crescator
	_ = []int{1123, 4151, 2, 536436, 8, 123}
	var data sort.Interface
	sort.Sort(data)
	fmt.Println()
}

func TestSorting(t *testing.T) {
	testData := []struct {
		name   string
		in     []int
		expect []int
	}{
		{
			name:   "test de la Calin",
			in:     []int{1, 4, 3, 4, 2, 5},
			expect: []int{1, 3, 5, 2, 4, 4},
		},
		{
			name:   "test de la SRP",
			in:     []int{1123, 4151, 2, 536436, 8, 123},
			expect: []int{123, 1123, 4151, 2, 8, 536436},
		},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			data := interface_implementations.IntSliceSort(td.in)
			sort.Sort(data)
			require.Equal(t, td.expect, td.in)
		})
	}
}

type CorinaSorter struct {
	ints []int
}

func (ss CorinaSorter) Len() int {
	return len(ss.ints)
}

func (ss CorinaSorter) Less(i, j int) bool {
	if ss.ints[i]%2 == 1 && ss.ints[j]%2 == 0 { // daca prima este impara, a doua para
		return true //
	}

	if ss.ints[i]%2 == ss.ints[j]%2 {
		return ss.ints[i] < ss.ints[j] // comparata egalitatea paritatilor
	}

	return false // prima e para
}

func (ss CorinaSorter) Swap(i, j int) {
	ss.ints[i], ss.ints[j] = ss.ints[j], ss.ints[i]
}
