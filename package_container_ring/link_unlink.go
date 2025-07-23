package package_container_ring

import (
	"container/ring"
	"fmt"
)

func LinkUnlinkRing() {
	r1 := ring.New(2)
	r2 := ring.New(2)

	r1.Value = "A"
	r1.Next().Value = "B"

	r2.Value = "X"
	r2.Next().Value = "Y"

	// Gabungkan r2 ke r1
	r1.Link(r2)

	// Tampilkan semua isi
	r1.Do(func(p any) {
		fmt.Print(p, " ")
	})
	fmt.Println()

	// Unlink 2 elemen setelah r1
	r1.Unlink(2)

	// Tampilkan isi setelah unlink
	r1.Do(func(p any) {
		fmt.Print(p, " ")
	})
	fmt.Println()
}
