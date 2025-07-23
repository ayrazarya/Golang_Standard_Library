package package_fmt

import "fmt"

func DemoSprint() {
	fmt.Println("=== Sprint / Sprintf / Sprintln ===")

	nama := "Nugroho"
	umur := 25

	// Menggabungkan ke string tanpa cetak langsung
	result := fmt.Sprintf("Nama: %s | Umur: %d", nama, umur)
	fmt.Println("Hasil Sprintf:", result)

	// Sprint membuat string dari beberapa argumen
	s := fmt.Sprint("Halo", " ", nama)
	fmt.Println("Sprint:", s)
}
