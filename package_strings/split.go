package package_strings

import (
	"fmt"
	"strings"
)

func DemoSplit() {
	fmt.Println("\n=== Demo Split ===")

	text := "apel,jeruk,pisang"
	slice := strings.Split(text, ",")
	for i, fruit := range slice {
		fmt.Printf("Buah %d: %s\n", i+1, fruit)
	}
}
