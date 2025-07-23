package package_reflect

import (
	"fmt"
	"reflect"
)

func DemoTypeCheck() {
	vals := []interface{}{42, "halo", true}

	for _, val := range vals {
		t := reflect.TypeOf(val)
		fmt.Printf("Tipe dari %v adalah %s\n", val, t.Name())
	}
}
