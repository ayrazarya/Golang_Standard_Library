package package_regexp

import (
	"fmt"
	"regexp"
)

// DemoReplace mendemokan ReplaceAll, ReplaceAllString, ReplaceAllFunc, ReplaceAllStringFunc
func DemoReplace() {
	pattern := `go[a-z]+`
	text := "golang is awesome, gofood, gopher, gone, go!"
	re := regexp.MustCompile(pattern)

	fmt.Println("ReplaceAll:", string(re.ReplaceAll([]byte(text), []byte("REPLACED"))))
	fmt.Println("ReplaceAllString:", re.ReplaceAllString(text, "REPLACED"))
	fmt.Println("ReplaceAllFunc:", string(re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		return []byte("FUNC")
	})))
	fmt.Println("ReplaceAllStringFunc:", re.ReplaceAllStringFunc(text, func(s string) string {
		return "FUNCSTR"
	}))
}
