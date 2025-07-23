package package_strings

import (
	"fmt"
	"strings"
)

func DemoFields() {
	fmt.Println("\n=== Demo Fields ===")

	sentence := "Belajar bahasa Go sangat menyenangkan"
	words := strings.Fields(sentence)
	for i, word := range words {
		fmt.Printf("Word %d: %s\n", i+1, word)
	}
}
