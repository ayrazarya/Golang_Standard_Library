package package_path

import (
	"fmt"
	"path"
)

func DemoBase() {
	p := "/a/b/c.txt"
	fmt.Println("Base:", path.Base(p))
}
