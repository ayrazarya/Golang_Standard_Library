package package_math

import (
	"fmt"
	"math"
)

func DemoMod() {
	fmt.Println("\n=== Demo math.Mod ===")

	x, y := 10.5, 3.2
	fmt.Printf("Mod(%.2f, %.2f) = %.2f\n", x, y, math.Mod(x, y))
}
