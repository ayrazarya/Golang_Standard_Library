package package_flag

import (
	"flag"
	"fmt"
	"os"
)

func DemoCustomUsage() {
	fmt.Println("\n=== Demo Custom Usage ===")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Penggunaan custom:")
		fmt.Fprintln(os.Stderr, "  -nama string : Nama user")
		fmt.Fprintln(os.Stderr, "  -umur int    : Umur user")
	}

	nama := flag.String("nama", "", "Nama user")
	umur := flag.Int("umur", 0, "Umur user")
	flag.Parse()

	fmt.Printf("Nama: %s, Umur: %d\n", *nama, *umur)
}
