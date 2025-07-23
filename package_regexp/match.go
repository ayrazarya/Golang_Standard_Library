package package_regexp

import (
	"fmt"
	"regexp"
	"strings"
)

// DemoMatch mendemokan Match, MatchString, MatchReader
func DemoMatch() {
	pattern := `go+`
	data := []byte("gopher gooo go gooooolang")
	matched, err := regexp.Match(pattern, data)
	fmt.Println("regexp.Match:", matched, err)

	matchedStr, err := regexp.MatchString(pattern, "gooooolang")
	fmt.Println("regexp.MatchString:", matchedStr, err)

	re := regexp.MustCompile(pattern)
	matchedReader := re.MatchReader(strings.NewReader("gooooolang"))
	fmt.Println("regexp.MatchReader:", matchedReader)
}
