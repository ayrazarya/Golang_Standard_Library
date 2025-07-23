package package_regexp

import (
	"fmt"
	"regexp"
)

// DemoCompile mendemokan Compile dan MustCompile
func DemoCompile() {
	pattern := `go+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Compile error:", err)
		return
	}
	fmt.Println("regexp.Compile berhasil:", re.MatchString("gooooolang"))

	re2 := regexp.MustCompile(pattern)
	fmt.Println("regexp.MustCompile berhasil:", re2.MatchString("gooo"))
}

// DemoCompilePOSIX mendemokan CompilePOSIX dan MustCompilePOSIX
func DemoCompilePOSIX() {
	pattern := `go+`
	re, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		fmt.Println("CompilePOSIX error:", err)
		return
	}
	fmt.Println("regexp.CompilePOSIX berhasil:", re.MatchString("gooooolang"))

	re2 := regexp.MustCompilePOSIX(pattern)
	fmt.Println("regexp.MustCompilePOSIX berhasil:", re2.MatchString("gooo"))
}
