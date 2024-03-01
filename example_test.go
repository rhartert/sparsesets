package sparsesets

import "fmt"

func ExampleSet_Insert() {
	ss := New(5)
	ss.Insert(2)
	ss.Insert(4)

	fmt.Println(ss.String())
	// Output: {0: false, 1: false, 2: true, 3: false, 4: true}
}

func ExampleSet_Content() {
	ss := New(5)

	// Perform some mutations, inserting and removing elements.
	ss.Insert(2)
	ss.Insert(3)
	ss.Insert(4)
	ss.Remove(3)

	// Iterate on the current content of the set.
	for _, e := range ss.Content() {
		fmt.Println(e)
	}
	// Output:
	// 2
	// 4
}

func ExampleSet_Absent() {
	ss := New(5)

	// Perform some mutations, inserting and removing elements.
	ss.Insert(2)
	ss.Insert(3)
	ss.Insert(4)
	ss.Remove(3)

	// Iterate on the elements currently NOT in the set. Note that no guarantee
	// is provided on the order in which elements are returned.
	for _, e := range ss.Absent() {
		fmt.Printf("%d: %v\n", e, ss.Contains(e))
	}
	// Output:
	// 3: false
	// 1: false
	// 0: false
}
