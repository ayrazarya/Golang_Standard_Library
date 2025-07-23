package package_encoding

import (
	"encoding/csv"
	"fmt"
	"strings"
)

// DemoCSV mendemokan parsing dan encoding CSV
func DemoCSV() {
	csvData := "nama,umur\nBudi,25\nAndi,30"
	r := csv.NewReader(strings.NewReader(csvData))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("CSV Read error:", err)
		return
	}
	fmt.Println("CSV Parsed:", records)

	var sb strings.Builder
	w := csv.NewWriter(&sb)
	w.WriteAll([][]string{{"nama", "umur"}, {"Cici", "22"}, {"Dodi", "28"}})
	fmt.Println("CSV Encoded:", sb.String())
}
