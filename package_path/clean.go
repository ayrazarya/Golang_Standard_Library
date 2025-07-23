package package_path

import (
	"fmt"
	"path"
)

func DemoClean() {
	fmt.Println("Clean:", path.Clean("/a//b/../c/./d/"))
}
