package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoParseFloat() {
	fmt.Println("\n=== Demo ParseFloat ===")

	s := "3.1415"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("string: %q → float64: %f\n", s, f)
	}
}
