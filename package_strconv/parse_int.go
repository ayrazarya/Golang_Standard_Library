package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoParseInt() {
	fmt.Println("\n=== Demo ParseInt ===")

	s := "42"
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("string: %q → int64: %d\n", s, i)
	}
}
