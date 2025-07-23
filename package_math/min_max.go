package package_math

import (
	"fmt"
	"math"
)

func DemoMinMax() {
	fmt.Println("\n=== Demo math.Min & math.Max ===")

	a, b := 4.2, 7.9
	fmt.Printf("Min(%.2f, %.2f) = %.2f\n", a, b, math.Min(a, b))
	fmt.Printf("Max(%.2f, %.2f) = %.2f\n", a, b, math.Max(a, b))
}
