package package_path

import (
	"fmt"
	"path"
)

func DemoMatch() {
	matched, err := path.Match("*.txt", "c.txt")
	fmt.Println("Match '*.txt' dengan 'c.txt':", matched, err)
}
