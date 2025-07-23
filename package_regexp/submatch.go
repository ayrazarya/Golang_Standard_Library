package package_regexp

import (
	"fmt"
	"regexp"
)

// DemoSubmatch mendemokan FindSubmatch, FindAllSubmatch, FindSubmatchIndex, FindAllSubmatchIndex, FindStringSubmatch, FindAllStringSubmatch, FindStringSubmatchIndex, FindAllStringSubmatchIndex
func DemoSubmatch() {
	pattern := `go([a-z]+)`
	text := "golang is awesome, gofood, gopher, gone, go!"
	re := regexp.MustCompile(pattern)
	fmt.Println("FindSubmatch:", re.FindSubmatch([]byte(text)))
	fmt.Println("FindAllSubmatch:", re.FindAllSubmatch([]byte(text), -1))
	fmt.Println("FindSubmatchIndex:", re.FindSubmatchIndex([]byte(text)))
	fmt.Println("FindAllSubmatchIndex:", re.FindAllSubmatchIndex([]byte(text), -1))
	fmt.Println("FindStringSubmatch:", re.FindStringSubmatch(text))
	fmt.Println("FindAllStringSubmatch:", re.FindAllStringSubmatch(text, -1))
	fmt.Println("FindStringSubmatchIndex:", re.FindStringSubmatchIndex(text))
	fmt.Println("FindAllStringSubmatchIndex:", re.FindAllStringSubmatchIndex(text, -1))
}
