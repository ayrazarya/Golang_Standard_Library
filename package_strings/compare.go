package package_strings

import (
	"fmt"
	"strings"
)

func DemoCompare() {
	fmt.Println("\n=== Demo Compare ===")

	fmt.Println("Compare('a', 'b'):", strings.Compare("a", "b"))
	fmt.Println("Compare('b', 'a'):", strings.Compare("b", "a"))
	fmt.Println("Compare('a', 'a'):", strings.Compare("a", "a"))
}
