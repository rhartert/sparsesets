// Package sparsesets provides an implementation of the sparse set data
// structure.
//
// A sparse set is an optimized data structure to perform efficient operations
// on sets of unique integers from 0 to N-1. It achieves optimal time complexity
// for most of its operations:
//   - Insert:   O(1)
//   - Remove:   O(1)
//   - Contains: O(1)
//   - Clear:    O(1)
//   - Content:  O(size)
//   - Absent:   O(N - size)
//
// This efficiency comes at the cost of increased memory usage, as a sparse set
// internally utilizes two slices of N integers.
package sparsesets

import (
	"fmt"
	"strings"
)

type Set struct {
	size      int
	values    []int
	positions []int
}

// New creates and initializes a new empty Set for elements from 0 to N-1.
//
// Example:
//
//	set := New(10) // creates a Set to hold integers [0, 9]
//
// The function panics if n is negative.
func New(n int) *Set {
	if n < 0 {
		panic(fmt.Sprintf("negative capacity n: %d", n))
	}
	ss := &Set{
		size:      0,
		values:    make([]int, n),
		positions: make([]int, n),
	}
	for i := 0; i < n; i++ {
		ss.values[i] = i
		ss.positions[i] = i
	}
	return ss
}

// Contains returns true if elem is in the set; it returns false otherwise.
func (ss *Set) Contains(elem int) bool {
	return ss.positions[elem] < ss.size
}

// N returns the capacity of the sparse set, which represents the range of
// possible values (0 to N-1) that the set can contain.
func (ss *Set) N() int {
	return len(ss.values)
}

// Size returns the number of elements contained in the sparse set.
func (ss *Set) Size() int {
	return ss.size
}

// Remove removes elem from the set. It returns true if the element was in the
// set; false otherwise.
func (ss *Set) Remove(elem int) error {
	if elem >= ss.N() {
		return fmt.Errorf("elem %d is out of the set range [0, %d)", elem, ss.N())
	}
	p := ss.positions[elem]
	ss.size -= 1
	if p >= ss.size { // elem is already out of the set
		return nil
	}
	ss.swap(p, ss.size)
	return nil
}

// Insert insets elem in the set. It returns true if the element was not in the
// set; false otherwise.
func (ss *Set) Insert(elem int) error {
	if elem >= ss.N() {
		return fmt.Errorf("elem %d is out of the set range [0, %d)", elem, ss.N())
	}
	p := ss.positions[elem]
	if p < ss.size { // elem is already in the set
		return nil
	}
	ss.swap(p, ss.size)
	ss.size += 1
	return nil

}

func (ss *Set) swap(p1, p2 int) {
	if p1 == p2 {
		return
	}
	v1 := ss.values[p1]
	v2 := ss.values[p2]
	ss.values[p1] = v2
	ss.values[p2] = v1
	ss.positions[v1] = p2
	ss.positions[v2] = p1
}

// Clear empties the set in constant time.
func (ss *Set) Clear() {
	ss.size = 0
}

// Content returns a slice containing the current elements of the set.
//
// Important: The returned slice should only be used for read-only operations as
// it is a direct view of the set's internal structure. Modifications to this
// slice can lead to undefined behavior of the set. Furthermore, any mutation
// to the set (for instance, by calling Insert or Remove) may invalidate the
// contents of the slice, rendering it unreliable.
func (ss *Set) Content() []int {
	return ss.values[:ss.size]
}

// Absent returns a slice containing the elements currently absent from the set.
//
// Important: The returned slice should only be used for read-only operations as
// it is a direct view of the set's internal structure. Modifications to this
// slice can lead to undefined behavior of the set. Furthermore, any mutation
// to the set (for instance, by calling Insert or Remove) may invalidate the
// contents of the slice, rendering it unreliable.
func (ss *Set) Absent() []int {
	return ss.values[ss.size:]
}

// String returns a string representation of the Set, formatting the set's
// elements and their presence in a map-like structure. Each element from
// 0 to N-1 is followed by a boolean indicating whether it is present (true)
// or absent (false) in the set.
func (ss *Set) String() string {
	if len(ss.positions) == 0 {
		return "{}"
	}
	bf := strings.Builder{}
	bf.WriteString("{")
	for v, p := range ss.positions[:len(ss.positions)-1] {
		bf.WriteString(fmt.Sprintf("%d: %v, ", v, p < ss.size))
	}
	lv := len(ss.positions) - 1
	lp := ss.positions[lv]
	bf.WriteString(fmt.Sprintf("%d: %v}", lv, lp < ss.size))
	return bf.String()
}
