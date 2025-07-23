package package_strings

import (
	"fmt"
	"strings"
)

func DemoHasPrefixSuffix() {
	fmt.Println("\n=== Demo HasPrefix & HasSuffix ===")

	s := "main.go"
	fmt.Println("HasPrefix 'main'? :", strings.HasPrefix(s, "main"))
	fmt.Println("HasSuffix '.go'?  :", strings.HasSuffix(s, ".go"))
}
