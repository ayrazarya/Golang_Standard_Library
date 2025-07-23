package package_path

import (
	"fmt"
	"path"
)

func DemoDir() {
	p := "/a/b/c.txt"
	fmt.Println("Dir:", path.Dir(p))
}
