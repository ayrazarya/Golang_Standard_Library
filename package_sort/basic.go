package package_sort

import (
	"fmt"
	"sort"
)

func BasicSort() {
	ints := []int{5, 2, 4, 3, 1}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	strings := []string{"banana", "apple", "cherry"}
	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings)
}
