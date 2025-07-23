package package_os

import (
	"fmt"
	"os"
)

func DemoCreateFileAndDir() {
	fmt.Println("\n=== Demo Create File & Directory ===")

	// Buat folder
	err := os.MkdirAll("demo_folder", 0755)
	if err != nil {
		fmt.Println("Gagal membuat folder:", err)
		return
	}

	// Buat file di dalamnya
	file, err := os.Create("demo_folder/hello.txt")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Halo dari file baru!\n")
	fmt.Println("Folder dan file berhasil dibuat.")
}
