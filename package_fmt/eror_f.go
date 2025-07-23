package package_fmt

import (
	"errors"
	"fmt"
)

func DemoErrorf() {
	fmt.Println("\n=== Demo Errorf ===")

	// Error biasa
	err := errors.New("file tidak ditemukan")

	// Membungkus error dengan Errorf
	wrapped := fmt.Errorf("gagal membuka file: %w", err)

	fmt.Println("Error asli:", err)
	fmt.Println("Error wrapped:", wrapped)
}
