package package_math

import (
	"fmt"
	"math"
)

func DemoLog() {
	fmt.Println("\n=== Demo Logaritma ===")

	x := 10.0
	fmt.Printf("Log(%.2f) base e = %.2f\n", x, math.Log(x))
	fmt.Printf("Log10(%.2f)      = %.2f\n", x, math.Log10(x))
}
