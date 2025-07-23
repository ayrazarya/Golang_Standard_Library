package package_fmt

import (
	"fmt"
)

func DemoScan() {
	fmt.Println("\n=== Demo: Scan / Scanf / Scanln ===")

	// -------------------------------
	// 1. fmt.Scan()
	// Membaca input berurutan (dipisah spasi), misalnya: "Arya 21"
	// -------------------------------
	var nama string
	var umur int

	fmt.Print("Masukkan nama dan umur (dipisah spasi): ")
	fmt.Scan(&nama, &umur)

	fmt.Printf("Halo %s, umurmu %d\n", nama, umur)

	// -------------------------------
	// 2. fmt.Scanln()
	// Sama seperti Scan tapi hanya sampai newline (ENTER)
	// Cocok untuk input satu baris
	// -------------------------------
	var hobi string
	fmt.Print("Masukkan hobi kamu: ")
	fmt.Scanln(&hobi)

	fmt.Println("Hobimu:", hobi)

	// -------------------------------
	// 3. fmt.Scanf()
	// Input dengan format tertentu
	// -------------------------------
	var kota string
	fmt.Print("Masukkan kota asal (format: Kota: <nama_kota>): ")
	fmt.Scanf("Kota: %s", &kota)

	fmt.Println("Kota asal:", kota)
}
