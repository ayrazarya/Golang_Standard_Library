package package_io

import (
	"bufio"
	"fmt"
	"strings"
)

// DemoBufio mendemokan bufio.Reader, bufio.Writer, bufio.Scanner
func DemoBufio() {
	// Reader
	reader := bufio.NewReader(strings.NewReader("halo\ndunia\n"))
	line, _ := reader.ReadString('\n')
	fmt.Print("ReadString: ", line)
	// Scanner
	scanner := bufio.NewScanner(strings.NewReader("baris1\nbaris2\n"))
	for scanner.Scan() {
		fmt.Println("Scanner:", scanner.Text())
	}
	// Writer
	var sb strings.Builder
	writer := bufio.NewWriter(&sb)
	writer.WriteString("tulis ke buffer\n")
	writer.Flush()
	fmt.Print("Writer: ", sb.String())
}
