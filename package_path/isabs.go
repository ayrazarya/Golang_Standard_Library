package package_path

import (
	"fmt"
	"path"
)

func DemoIsAbs() {
	fmt.Println("IsAbs '/a/b':", path.IsAbs("/a/b"))
	fmt.Println("IsAbs 'a/b':", path.IsAbs("a/b"))
}
