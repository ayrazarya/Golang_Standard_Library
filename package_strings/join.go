package package_strings

import (
	"fmt"
	"strings"
)

func DemoJoin() {
	fmt.Println("\n=== Demo Join ===")

	words := []string{"go", "is", "awesome"}
	result := strings.Join(words, " ")
	fmt.Println("Join:", result)
}
