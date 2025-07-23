package package_slices

import (
	"fmt"
	"slices"
)

// DemoReverseReplace mendemokan Reverse dan Replace
func DemoReverseReplace() {
	s := []int{1, 2, 3, 4}
	slices.Reverse(s)
	fmt.Println("Reverse:", s)
	s2 := []int{5, 6}
	slices.Replace(s, 1, 3, s2...)
	fmt.Println("Replace index 1-2 dengan [5 6]:", s)
}
