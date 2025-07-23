package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoLengthAndAccess() {
	fmt.Println("\n=== Demo Len(), Front(), Back() ===")

	l := list.New()
	l.PushBack("A")
	l.PushBack("B")
	l.PushBack("C")

	fmt.Println("Panjang list:", l.Len())
	fmt.Println("Depan:", l.Front().Value)
	fmt.Println("Belakang:", l.Back().Value)
}
