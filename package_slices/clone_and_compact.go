package package_slices

import (
	"fmt"
	"slices"
)

// DemoCloneCompact mendemokan Clone dan Compact
func DemoCloneCompact() {
	s := []int{1, 2, 2, 3, 3, 3, 4}
	clone := slices.Clone(s)
	fmt.Println("Clone:", clone)
	compacted := slices.Compact(clone)
	fmt.Println("Compact:", compacted)
}
