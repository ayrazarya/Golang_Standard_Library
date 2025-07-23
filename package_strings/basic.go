package package_strings

import (
	"fmt"
	"strings"
)

func DemoBasic() {
	fmt.Println("\n=== Demo Basic ===")

	text := "Hello Golang"
	fmt.Println("Original :", text)
	fmt.Println("ToLower  :", strings.ToLower(text))
	fmt.Println("ToUpper  :", strings.ToUpper(text))
}
