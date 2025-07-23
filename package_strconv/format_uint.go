package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoFormatUint() {
	fmt.Println("\n=== Demo FormatUint ===")

	u := uint64(123)
	s := strconv.FormatUint(u, 10)
	fmt.Printf("uint64: %d → string: %q\n", u, s)
}
