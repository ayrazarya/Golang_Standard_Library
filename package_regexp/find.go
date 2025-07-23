package package_regexp

import (
	"fmt"
	"regexp"
)

// DemoFind mendemokan Find, FindAll, FindIndex, FindAllIndex, FindString, FindAllString, FindStringIndex, FindAllStringIndex
func DemoFind() {
	pattern := `go[a-z]+`
	text := "golang is awesome, gofood, gopher, gone, go!"
	re := regexp.MustCompile(pattern)
	fmt.Println("Find:", string(re.Find([]byte(text))))
	fmt.Println("FindAll:", re.FindAll([]byte(text), -1))
	fmt.Println("FindIndex:", re.FindIndex([]byte(text)))
	fmt.Println("FindAllIndex:", re.FindAllIndex([]byte(text), -1))
	fmt.Println("FindString:", re.FindString(text))
	fmt.Println("FindAllString:", re.FindAllString(text, -1))
	fmt.Println("FindStringIndex:", re.FindStringIndex(text))
	fmt.Println("FindAllStringIndex:", re.FindAllStringIndex(text, -1))
}
