package package_strings

import (
	"fmt"
	"strings"
)

func DemoEqualFold() {
	fmt.Println("\n=== Demo EqualFold (case-insensitive compare) ===")

	fmt.Println("EqualFold('Go', 'go'):", strings.EqualFold("Go", "go"))
	fmt.Println("EqualFold('Hello', 'HELLO'):", strings.EqualFold("Hello", "HELLO"))
}
