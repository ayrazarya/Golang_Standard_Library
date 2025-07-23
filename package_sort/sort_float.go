package package_sort

import (
	"fmt"
	"sort"
)

func SortFloat() {
	data := []float64{4.1, 2.9, 3.3, 1.5}
	sort.Float64s(data)
	fmt.Println("Sorted floats:", data)
}
