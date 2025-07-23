package package_os

import (
	"fmt"
	"os"
)

func DemoEnv() {
	fmt.Println("\n=== Demo Environment Variable ===")

	user := os.Getenv("USER") // Linux/Mac
	if user == "" {
		user = os.Getenv("USERNAME") // Windows
	}
	fmt.Println("User login:", user)

	// Set env (tidak permanen)
	os.Setenv("GOLANG_MODE", "development")
	fmt.Println("GOLANG_MODE:", os.Getenv("GOLANG_MODE"))
}
