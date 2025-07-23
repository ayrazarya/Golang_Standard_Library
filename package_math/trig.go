package package_math

import (
	"fmt"
	"math"
)

func DemoTrig() {
	fmt.Println("\n=== Demo Trigonometri (math.Sin, math.Cos, math.Tan) ===")

	angle := math.Pi / 4 // 45 derajat
	fmt.Printf("Sin(π/4) = %.2f\n", math.Sin(angle))
	fmt.Printf("Cos(π/4) = %.2f\n", math.Cos(angle))
	fmt.Printf("Tan(π/4) = %.2f\n", math.Tan(angle))
}
