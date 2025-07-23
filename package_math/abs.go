package package_math

import (
	"fmt"
	"math"
)

func DemoAbs() {
	fmt.Println("\n=== Demo math.Abs ===")

	n := -5.7
	fmt.Printf("Nilai absolut dari %.2f = %.2f\n", n, math.Abs(n))
}
