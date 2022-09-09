package set_test

import (
	"fmt"

	"github.com/micnncim/go-set"
)

func ExampleSetHas() {
	s := set.New(1, 2)

	fmt.Println(s.Has(1))
	fmt.Println(s.Has(0))
	// Output:
	// true
	// false
}

func ExampleSetHasAll() {
	s := set.New("foo", "bar")

	fmt.Println(s.HasAll("foo", "bar"))
	fmt.Println(s.HasAll("foo", "bar", "baz"))
	// Output:
	// true
	// false
}

func ExampleSetDifference() {
	s := set.New(1, 2)
	t := set.New(2, 3)

	fmt.Println(s.Difference(t))
	// Output:
	// [1]
}

func ExampleSetIntersection() {
	s := set.New(1, 2)
	t := set.New(2, 4)

	fmt.Println(s.Intersection(t))
	// Output:
	// [2]
}
