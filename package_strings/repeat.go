package package_strings

import (
	"fmt"
	"strings"
)

func DemoRepeat() {
	fmt.Println("\n=== Demo Repeat ===")

	word := "na"
	fmt.Println("Repeat 3x:", strings.Repeat(word, 3))
}
