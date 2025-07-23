package package_errors

import "fmt"

func Penjelasan() {
	fmt.Println("=== Modul Error di Go ===")
	fmt.Println("- Error di Go adalah interface dengan satu method: Error() string")
	fmt.Println("- Error umum dibuat dengan errors.New() atau fmt.Errorf()")
	fmt.Println("- Kita bisa buat custom error pakai struct")
	fmt.Println("- Error bisa dibungkus (wrapped) untuk trace asalnya")
	fmt.Println("- Type assertion berguna untuk deteksi jenis error\n")
}
