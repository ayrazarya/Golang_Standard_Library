package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoBasic() {
	fmt.Println("\n=== Demo Basic List ===")

	l := list.New()
	l.PushBack("A")
	l.PushBack("B")
	l.PushFront("Start")

	fmt.Println("Isi list:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("-", e.Value)
	}
}
