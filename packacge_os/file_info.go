package package_os

import (
	"fmt"
	"os"
)

func DemoFileInfo() {
	fmt.Println("\n=== Demo File Info (os.Stat) ===")

	info, err := os.Stat("sample.txt")
	if os.IsNotExist(err) {
		fmt.Println("File tidak ditemukan.")
		return
	}
	if err != nil {
		fmt.Println("Gagal mendapatkan info file:", err)
		return
	}

	fmt.Println("Nama  :", info.Name())
	fmt.Println("Ukuran:", info.Size())
	fmt.Println("Modus :", info.Mode())
	fmt.Println("Modifikasi terakhir:", info.ModTime())
}
