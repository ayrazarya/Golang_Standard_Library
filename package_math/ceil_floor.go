package package_math

import (
	"fmt"
	"math"
)

func DemoCeilFloor() {
	fmt.Println("\n=== Demo math.Ceil & math.Floor ===")

	val := 2.3
	fmt.Printf("Ceil(%.2f)  = %.2f\n", val, math.Ceil(val))
	fmt.Printf("Floor(%.2f) = %.2f\n", val, math.Floor(val))
}
