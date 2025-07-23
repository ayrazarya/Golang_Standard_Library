package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoParseUint() {
	fmt.Println("\n=== Demo ParseUint ===")

	s := "123"
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("string: %q → uint64: %d\n", s, u)
	}
}
