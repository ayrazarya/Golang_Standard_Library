package package_container_ring

import (
	"container/ring"
	"fmt"
)

func MoveRing() {
	r := ring.New(4)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// Bergerak maju 2 langkah
	r = r.Move(2)
	fmt.Println("Value setelah Move(2):", r.Value)

	// Bergerak mundur 1 langkah
	r = r.Move(-1)
	fmt.Println("Value setelah Move(-1):", r.Value)
}
