package package_container_ring

import (
	"container/ring"
	"fmt"
)

func BasicRing() {
	r := ring.New(5)

	// Mengisi ring dengan angka 1 sampai 5
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	// Menampilkan isi ring
	r.Do(func(p any) {
		fmt.Print(p, " ")
	})
	fmt.Println()
}
