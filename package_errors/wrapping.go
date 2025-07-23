package package_errors

import (
	"errors"
	"fmt"
)

var ErrDBNotFound = errors.New("database tidak ditemukan")

func DemoWrappingError() {
	fmt.Println("\n=== Demo Wrapping Error ===")

	// Buat error asli
	original := ErrDBNotFound

	// Bungkus error-nya
	wrapped := fmt.Errorf("query gagal: %w", original)

	fmt.Println("Original :", original)
	fmt.Println("Wrapped  :", wrapped)

	// Gunakan errors.Is untuk deteksi apakah wrapped mengandung original
	if errors.Is(wrapped, ErrDBNotFound) {
		fmt.Println("✔️ Wrapped error mengandung ErrDBNotFound (errors.Is)")
	} else {
		fmt.Println("❌ Wrapped error tidak mengandung ErrDBNotFound")
	}
}
