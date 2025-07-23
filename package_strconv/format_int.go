package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoFormatInt() {
	fmt.Println("\n=== Demo FormatInt ===")

	i := int64(42)
	s := strconv.FormatInt(i, 10)
	fmt.Printf("int64: %d → string: %q\n", i, s)
}
