package package_path

import (
	"fmt"
	"path"
)

func DemoExt() {
	p := "/a/b/c.txt"
	fmt.Println("Ext:", path.Ext(p))
}
