package package_fmt

import (
	"fmt"
	"os"
)

func DemoFprint() {
	fmt.Println("\n=== Demo Fprint / Fprintf / Fprintln ===")

	fmt.Fprint(os.Stdout, "Ini Fprint ke Stdout\n")

	nama := "Arya"
	umur := 22
	fmt.Fprintf(os.Stdout, "Nama: %s | Umur: %d\n", nama, umur)

	fmt.Fprintln(os.Stdout, "Ini dari Fprintln ke Stdout")
}
