package package_container_ring

import (
	"container/ring"
	"fmt"
)

func NextPrevRing() {
	r := ring.New(3)
	r.Value = "A"
	r.Next().Value = "B"
	r.Next().Next().Value = "C"

	fmt.Println("Current:", r.Value)
	fmt.Println("Next:", r.Next().Value)
	fmt.Println("Prev:", r.Prev().Value)
}
