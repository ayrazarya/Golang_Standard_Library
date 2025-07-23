package package_sort

import (
	"fmt"
	"sort"
)

func SortString() {
	data := []string{"zebra", "apple", "mango", "banana"}
	sort.Strings(data)
	fmt.Println("Sorted strings:", data)
}
