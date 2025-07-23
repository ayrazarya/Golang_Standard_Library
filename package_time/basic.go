package package_time

import (
	"fmt"
	"time"
)

func BasicUsage() {
	now := time.Now()
	fmt.Println("Waktu sekarang:", now)
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Format custom:", formatted)

	parsed, _ := time.Parse("2006-01-02", "2025-07-23")
	fmt.Println("Parsed date:", parsed)

	fmt.Println("Tahun:", now.Year())
	fmt.Println("Bulan:", now.Month())
	fmt.Println("Hari:", now.Day())
	fmt.Println("Jam:", now.Hour())
	fmt.Println("Menit:", now.Minute())
	fmt.Println("Detik:", now.Second())
}
