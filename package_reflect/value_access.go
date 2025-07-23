package package_reflect

import (
	"fmt"
	"reflect"
)

func DemoValueAccess() {
	num := 10
	v := reflect.ValueOf(&num).Elem()

	fmt.Println("Before:", v.Int())
	v.SetInt(99)
	fmt.Println("After:", v.Int())
	fmt.Println("Asli:", num)
}
