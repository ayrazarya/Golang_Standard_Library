package package_regexp

import (
	"fmt"
	"regexp"
)

// UsageRegexpDemo menampilkan beberapa fitur regexp
func UsageRegexpDemo() {
	pattern := `go([a-z]+)`
	text := "golang is awesome, gofood, gopher, gone, go!"
	re := regexp.MustCompile(pattern)

	fmt.Println("Semua match:", re.FindAllString(text, -1))
	fmt.Println("Apakah 'gopher' match?", re.MatchString("gopher"))
	fmt.Println("Replace 'go' dengan 'GO':", re.ReplaceAllString(text, "GO$1"))
}
