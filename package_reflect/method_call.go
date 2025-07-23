package package_reflect

import (
	"fmt"
	"reflect"
)

type Animal struct{}

func (a Animal) Speak(name string) {
	fmt.Println("Hello,", name)
}

func DemoMethodCall() {
	a := Animal{}
	method := reflect.ValueOf(a).MethodByName("Speak")
	args := []reflect.Value{reflect.ValueOf("Dino")}
	method.Call(args)
}
