package package_math

import (
	"fmt"
	"math"
)

func DemoPow() {
	fmt.Println("\n=== Demo math.Pow ===")

	a, b := 2.0, 3.0
	fmt.Printf("%.2f pangkat %.2f = %.2f\n", a, b, math.Pow(a, b))
}
