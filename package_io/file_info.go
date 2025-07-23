package package_io

import (
	"fmt"
	"os"
)

// DemoFileInfo mendemokan os.Stat, os.IsNotExist
func DemoFileInfo() {
	filename := "cek.txt"
	os.WriteFile(filename, []byte("cek"), 0644)
	info, err := os.Stat(filename)
	if err == nil {
		fmt.Println("Nama file:", info.Name())
		fmt.Println("Ukuran:", info.Size())
		fmt.Println("ModTime:", info.ModTime())
	}
	os.Remove(filename)
	_, err = os.Stat("tidakada.txt")
	fmt.Println("IsNotExist untuk file tidak ada:", os.IsNotExist(err))
}
