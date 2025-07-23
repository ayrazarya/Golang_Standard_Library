package package_os

import (
	"fmt"
	"os"
)

func DemoDelete() {
	fmt.Println("\n=== Demo Delete File/Folder ===")

	err := os.Remove("sample.txt")
	if err != nil {
		fmt.Println("Gagal menghapus file:", err)
	} else {
		fmt.Println("File sample.txt berhasil dihapus.")
	}

	err = os.RemoveAll("demo_folder")
	if err != nil {
		fmt.Println("Gagal menghapus folder:", err)
	} else {
		fmt.Println("Folder demo_folder berhasil dihapus.")
	}
}
