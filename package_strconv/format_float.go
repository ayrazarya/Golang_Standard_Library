package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoFormatFloat() {
	fmt.Println("\n=== Demo FormatFloat ===")

	f := 3.1415
	s := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("float64: %f → string: %q\n", f, s)
}
