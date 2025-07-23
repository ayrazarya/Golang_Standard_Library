package package_sort

import (
	"fmt"
	"sort"
)

func SortInt() {
	data := []int{9, 1, 6, 3, 7}
	sort.Ints(data)
	fmt.Println("Sorted ints:", data)
}
