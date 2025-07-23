package package_regexp

import (
	"fmt"
	"regexp"
)

// BasicRegexpDemo menampilkan penggunaan dasar regexp
func BasicRegexpDemo() {
	pattern := `a[a-z]+e`
	text := "apple avenue axe age ape"

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	fmt.Println("Pattern:", pattern)
	fmt.Println("Text:", text)
	fmt.Println("Matches:", matches)
}
