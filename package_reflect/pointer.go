package package_reflect

import (
	"fmt"
	"reflect"
)

func DemoPointerReflect() {
	var x int = 100
	ptr := reflect.ValueOf(&x)

	fmt.Println("Tipe:", ptr.Type())
	fmt.Println("Bisa di-set?", ptr.Elem().CanSet())

	ptr.Elem().SetInt(200)
	fmt.Println("Setelah diubah:", x)
}
