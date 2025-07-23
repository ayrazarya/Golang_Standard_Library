package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoAtoi() {
	fmt.Println("\n=== Demo Atoi (string to int) ===")

	s := "456"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("string: %q → int: %d\n", s, i)
	}
}
