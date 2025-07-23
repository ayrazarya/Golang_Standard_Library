package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoItoa() {
	fmt.Println("\n=== Demo Itoa (int to string) ===")

	n := 123
	s := strconv.Itoa(n)
	fmt.Printf("int: %d → string: %q\n", n, s)
}
