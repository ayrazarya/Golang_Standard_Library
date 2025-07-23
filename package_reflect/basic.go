package package_reflect

import (
	"fmt"
	"reflect"
)

func DemoBasicReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("Tipe:", v.Type())
	fmt.Println("Nilai:", v.Float())
	fmt.Println("Bisa diubah?", v.CanSet()) // false karena bukan pointer
}
