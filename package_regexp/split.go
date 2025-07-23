package package_regexp

import (
	"fmt"
	"regexp"
)

// DemoSplit mendemokan Split
func DemoSplit() {
	pattern := `,\s*`
	text := "apel, jeruk, pisang, mangga"
	re := regexp.MustCompile(pattern)
	result := re.Split(text, -1)
	fmt.Println("Split hasil:", result)
}
