package composition

import (
	"cmp"
	"esol/interface_implementations"
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"sort"
	"testing"
)

type CaiPutere int // []CaiPutere e sortabil direct

type Animal struct {
	Name string
	Age  int
}

func (a Animal) String() string {
	return fmt.Sprintf("La multean, %v! Ai facut %v.", a.Name, a.Age)
}

type Dog struct {
	Animal   Animal
	DogBreed string
}

func (a Dog) String() string {
	return fmt.Sprintf("La multean, %v! Ai facut %v.", a.Animal.Name, a.Animal.Age)
}

type DogWithComposition struct {
	Animal
	DogBreed string
}

//func (a DogWithComposition) String() string {
//	return fmt.Sprintf("La multean, %v! Ai facut %v. Esti cel mai dragut %v", a.Name, a.Age, a.DogBreed)
//}

func TestCompositionExample(t *testing.T) {
	someAnimal := Animal{
		Name: "Otto",
		Age:  3,
	}
	//otto := Dog{
	//	Animal:   someAnimal,
	//	DogBreed: "Bulldog Francez",
	//}

	ottoButAfterBeingFed := DogWithComposition{
		Animal:   someAnimal,
		DogBreed: "Bulldog Francez",
	}
	// otto.Name <- not possible

	fmt.Println("Animal", someAnimal)

	fmt.Println("CuCompositie", ottoButAfterBeingFed)
}

type Reverse struct {
	sort.Interface
}

func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func TestReverse(t *testing.T) {
	data := []int{1, 2, 3, 5, 4}
	sorter := interface_implementations.IntSliceSort(data)
	reversedSorter := Reverse{Interface: sorter}
	sort.Sort(reversedSorter)
	require.Equal(t, []int{4, 2, 5, 3, 1}, data)
}

// Example holds A1 and A2, Examples have a natural order: ExampleNaturalOrdering
type Example struct {
	A1, A2 string
}

var ExampleNaturalOrdering = func(a, b Example) int {
	if n := cmp.Compare(a.A1, b.A1); n != 0 {
		return n
	}
	return cmp.Compare(a.A2, b.A2)
}

func TestOrderingOnSlices(t *testing.T) {
	slices.Sort([]CaiPutere{})

	// Comparable
	slices.SortFunc([]Example{{A1: "A", A2: "A2"}, {A1: "A", A2: "A3"}}, ExampleNaturalOrdering)
}
