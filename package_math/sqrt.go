package package_math

import (
	"fmt"
	"math"
)

func DemoSqrt() {
	fmt.Println("\n=== Demo math.Sqrt ===")

	x := 16.0
	fmt.Printf("Akar dari %.2f = %.2f\n", x, math.Sqrt(x))
}
