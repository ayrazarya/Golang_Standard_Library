package package_errors

import (
	"errors"
	"fmt"
)

func DemoBasicError() {
	fmt.Println("=== Demo errors.New() dan fmt.Errorf() ===")

	err1 := errors.New("ini error standar")
	err2 := fmt.Errorf("error dengan format: %d + %d = %d", 2, 3, 5)

	fmt.Println("errors.New:", err1)
	fmt.Println("fmt.Errorf :", err2)
}
