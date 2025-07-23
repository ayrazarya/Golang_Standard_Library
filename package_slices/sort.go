package package_slices

import (
	"fmt"
	"slices"
)

// DemoSort mendemokan Sort, IsSorted, BinarySearch
func DemoSort() {
	s := []int{4, 2, 1, 3}
	slices.Sort(s)
	fmt.Println("Sorted:", s)
	fmt.Println("IsSorted:", slices.IsSorted(s))
	idx, found := slices.BinarySearch(s, 3)
	fmt.Println("BinarySearch 3:", idx, found)
}
