package package_errors

import "fmt"

// CustomError struct
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Kode %d: %s", e.Code, e.Message)
}

func DemoCustomError() {
	fmt.Println("\n=== Demo Custom Error ===")

	err := &CustomError{Code: 400, Message: "Data tidak valid"}

	fmt.Println("Custom Error:", err)
}
