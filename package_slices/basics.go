package package_slices

import (
	"fmt"
	"slices"
)

// DemoBasics mendemokan Contains, Index, Equal
func DemoBasics() {
	s := []int{1, 2, 3, 4}
	fmt.Println("Contains 3:", slices.Contains(s, 3))
	fmt.Println("Index 4:", slices.Index(s, 4))
	fmt.Println("Equal dengan [1 2 3 4]:", slices.Equal(s, []int{1, 2, 3, 4}))
}
