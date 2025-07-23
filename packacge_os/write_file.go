package package_os

import (
	"fmt"
	"os"
)

func DemoWriteFile() {
	fmt.Println("\n=== Demo Write File ===")

	content := []byte("Ini adalah isi dari file baru.")
	err := os.WriteFile("sample.txt", content, 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}
	fmt.Println("File berhasil ditulis.")
}
