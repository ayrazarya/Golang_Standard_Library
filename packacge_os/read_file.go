package package_os

import (
	"fmt"
	"os"
)

func DemoReadFile() {
	fmt.Println("\n=== Demo Read File ===")

	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}
	fmt.Println("Isi file:")
	fmt.Println(string(data))
}
