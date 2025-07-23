package package_strings

import (
	"fmt"
	"strings"
)

func DemoReplace() {
	fmt.Println("\n=== Demo Replace ===")

	text := "saya suka golang, golang itu keren"
	fmt.Println("Replace 'golang' -> 'go' :", strings.Replace(text, "golang", "go", -1))
}
