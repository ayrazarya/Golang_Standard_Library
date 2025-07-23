package package_sort

import (
	"fmt"
	"sort"
)

func ReverseSort() {
	data := []int{10, 5, 8, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println("Reverse sorted:", data)
}
