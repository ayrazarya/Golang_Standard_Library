package package_flag

import (
	"flag"
	"fmt"
)

func DemoBooleanFlag() {
	fmt.Println("\n=== Demo Boolean Flag ===")

	debug := flag.Bool("debug", false, "Aktifkan mode debug")
	flag.Parse()

	if *debug {
		fmt.Println("Debug mode AKTIF")
	} else {
		fmt.Println("Debug mode NONAKTIF")
	}
}
