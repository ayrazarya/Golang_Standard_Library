package package_container_ring

import (
	"container/ring"
	"fmt"
)

func ValueSetRing() {
	r := ring.New(3)

	// Set value
	r.Value = "First"
	r.Next().Value = "Second"
	r.Next().Next().Value = "Third"

	// Access all
	r.Do(func(p any) {
		fmt.Println("Isi:", p)
	})
}
