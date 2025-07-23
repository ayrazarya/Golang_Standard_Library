package package_flag

import "fmt"

func PenjelasanFlag() {
	fmt.Println("Package 'flag' di Go digunakan untuk parsing argument command-line.")
	fmt.Println("Contoh pemakaian:")
	fmt.Println("  -name string  → Nama pengguna")
	fmt.Println("  -debug bool   → Mode debug")
	fmt.Println("Gunakan flag.Parse() untuk mengambil nilainya.")
}
