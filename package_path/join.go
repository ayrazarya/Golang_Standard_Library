package package_path

import (
	"fmt"
	"path"
)

func DemoJoin() {
	fmt.Println("Join:", path.Join("a", "b", "c.txt"))
}
