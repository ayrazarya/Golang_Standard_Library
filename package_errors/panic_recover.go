package package_errors

import "fmt"

func DemoPanicRecover() {
	fmt.Println("\n=== Demo Panic & Recover ===")
	safeFunction()
	fmt.Println("Program tetap lanjut setelah recover.")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("✅ Recovered dari panic:", r)
		}
	}()

	fmt.Println("Sebelum panic...")
	panic("🚨 Ini panic tiruan!")
	fmt.Println("Setelah panic (tidak akan dieksekusi)")
}
