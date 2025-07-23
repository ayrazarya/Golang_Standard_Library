package package_flag

import (
	"flag"
	"fmt"
)

func DemoDefaultValues() {
	fmt.Println("\n=== Demo Default Values ===")

	country := flag.String("country", "Indonesia", "Negara asal")
	flag.Parse()

	fmt.Println("Negara:", *country)
}
