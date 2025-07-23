package package_reflect

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func DemoStructInfo() {
	p := Person{Name: "Arya", Age: 22}
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)
	}
}
