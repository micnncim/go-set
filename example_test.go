package set_test

import (
	"fmt"

	"github.com/micnncim/go-set"
)

func ExampleSet_Has() {
	s := set.New(1, 2)

	fmt.Println(s.Contains(1))
	fmt.Println(s.Contains(0))
	// Output:
	// true
	// false
}

func ExampleSet_HasAll() {
	s := set.New("foo", "bar")

	fmt.Println(s.ContainsAll("foo", "bar"))
	fmt.Println(s.ContainsAll("foo", "bar", "baz"))
	// Output:
	// true
	// false
}

func ExampleSet_Difference() {
	s := set.New(1, 2)
	t := set.New(2, 3)

	fmt.Println(s.Difference(t))
	fmt.Println(t.Difference(s))
	// Output:
	// [1]
	// [3]
}

func ExampleSet_Intersection() {
	s := set.New(1, 2)
	t := set.New(2, 4)

	fmt.Println(s.Intersection(t))
	// Output:
	// [2]
}
