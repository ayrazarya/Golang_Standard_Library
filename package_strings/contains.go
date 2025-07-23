package package_strings

import (
	"fmt"
	"strings"
)

func DemoContains() {
	fmt.Println("\n=== Demo Contains ===")

	s := "Belajar Golang"
	fmt.Println("Contains 'Golang'? :", strings.Contains(s, "Golang"))
	fmt.Println("Contains 'Python'? :", strings.Contains(s, "Python"))
}
