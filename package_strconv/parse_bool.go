package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoParseBool() {
	fmt.Println("\n=== Demo ParseBool ===")

	inputs := []string{"true", "false", "1", "0"}
	for _, val := range inputs {
		b, err := strconv.ParseBool(val)
		if err != nil {
			fmt.Printf("'%s' → error: %v\n", val, err)
		} else {
			fmt.Printf("'%s' → bool: %t\n", val, b)
		}
	}
}
