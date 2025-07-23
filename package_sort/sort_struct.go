package package_sort

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func SortStruct() {
	people := []Person{
		{"Budi", 32},
		{"Ani", 25},
		{"Citra", 29},
	}
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:", people)
}
