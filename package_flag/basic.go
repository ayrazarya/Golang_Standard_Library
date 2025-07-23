package package_flag

import (
	"flag"
	"fmt"
)

func DemoBasicFlag() {
	fmt.Println("\n=== Demo Basic Flag ===")

	name := flag.String("name", "guest", "Nama pengguna")
	flag.Parse()

	fmt.Println("Halo,", *name)
}
