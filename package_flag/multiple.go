package package_flag

import (
	"flag"
	"fmt"
)

func DemoMultipleFlags() {
	fmt.Println("\n=== Demo Multiple Flags ===")

	host := flag.String("host", "localhost", "Alamat host")
	port := flag.Int("port", 8080, "Nomor port")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()

	fmt.Printf("Server berjalan di %s:%d\n", *host, *port)
	if *verbose {
		fmt.Println("Verbose mode aktif")
	}
}
