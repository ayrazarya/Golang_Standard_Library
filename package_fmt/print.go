package package_fmt

import "fmt"

func DemoPrint() {
	fmt.Println("=== Print / Println / Printf ===")
	fmt.Print("Ini Print tanpa newline. ")
	fmt.Print("Masih di baris yang sama.\n")

	fmt.Println("Ini Println dengan newline.")
	nama := "Arya"
	umur := 21
	fmt.Printf("Nama saya %s, umur saya %d\n", nama, umur)
}
