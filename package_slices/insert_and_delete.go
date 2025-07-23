package package_slices

import (
	"fmt"
	"slices"
)

// DemoInsertDelete mendemokan Insert dan Delete
func DemoInsertDelete() {
	s := []int{1, 2, 4, 5}
	s = slices.Insert(s, 2, 3)
	fmt.Println("Insert 3 di index 2:", s)
	s = slices.Delete(s, 1, 3)
	fmt.Println("Delete index 1-2:", s)
}
