package package_io

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// DemoFileRead mendemokan membaca file dengan berbagai cara
func DemoFileRead() {
	filename := "contoh.txt"
	os.WriteFile(filename, []byte("baris1\nbaris2\n"), 0644)
	defer os.Remove(filename)

	// ioutil.ReadFile
	data, _ := ioutil.ReadFile(filename)
	fmt.Print("ReadFile: ", string(data))

	// bufio.NewReader
	f, _ := os.Open(filename)
	defer f.Close()
	reader := bufio.NewReader(f)
	line, _ := reader.ReadString('\n')
	fmt.Print("bufio.ReadString: ", line)
}
