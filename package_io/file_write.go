package package_io

import (
	"bufio"
	"fmt"
	"os"
)

// DemoFileWrite mendemokan menulis file dengan berbagai cara
func DemoFileWrite() {
	filename := "tulis.txt"
	os.Remove(filename)
	// os.WriteFile
	os.WriteFile(filename, []byte("hello\n"), 0644)
	// bufio.NewWriter
	f, _ := os.Create(filename)
	writer := bufio.NewWriter(f)
	writer.WriteString("baris bufio\n")
	writer.Flush()
	f.Close()
	// Baca hasil
	data, _ := os.ReadFile(filename)
	fmt.Print("Isi file: ", string(data))
	os.Remove(filename)
}
