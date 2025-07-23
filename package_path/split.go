package package_path

import (
	"fmt"
	"path"
)

func DemoSplit() {
	p := "/a/b/c.txt"
	dir, file := path.Split(p)
	fmt.Println("Split dir:", dir, "file:", file)
}
