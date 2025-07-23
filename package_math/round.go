package package_math

import (
	"fmt"
	"math"
)

func DemoRound() {
	fmt.Println("\n=== Demo math.Round ===")

	fmt.Println("Round(2.3) =", math.Round(2.3))
	fmt.Println("Round(2.7) =", math.Round(2.7))
	fmt.Println("Round(-2.3) =", math.Round(-2.3))
}
