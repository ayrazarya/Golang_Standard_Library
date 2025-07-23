package package_io

import (
	"fmt"
	"os"
)

// DemoFileManip mendemokan os.Rename, os.Remove, os.Mkdir, os.MkdirAll
func DemoFileManip() {
	os.Mkdir("folder1", 0755)
	fmt.Println("Mkdir folder1")
	os.MkdirAll("folder2/sub", 0755)
	fmt.Println("MkdirAll folder2/sub")
	os.WriteFile("folder1/tes.txt", []byte("isi"), 0644)
	os.Rename("folder1/tes.txt", "folder1/tes2.txt")
	fmt.Println("Rename file di folder1")
	os.Remove("folder1/tes2.txt")
	fmt.Println("Remove file di folder1")
	os.RemoveAll("folder1")
	os.RemoveAll("folder2")
}
