package package_errors

import "fmt"

func DemoTypeAssertion() {
	fmt.Println("\n=== Demo Type Assertion terhadap Error ===")

	err := getError()

	if customErr, ok := err.(*CustomError); ok {
		fmt.Println("✔️ Ini CustomError dengan kode:", customErr.Code)
	} else {
		fmt.Println("❌ Bukan CustomError:", err)
	}
}

func getError() error {
	return &CustomError{Code: 403, Message: "Unauthorized access"}
}
