package package_strings

import (
	"fmt"
	"strings"
)

func DemoTrim() {
	fmt.Println("\n=== Demo Trim ===")

	text := "   halo dunia   "
	fmt.Println("TrimSpace:", strings.TrimSpace(text))
}
