package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoRemove() {
	fmt.Println("\n=== Demo Remove ===")

	l := list.New()
	e1 := l.PushBack("hapus saya")
	l.PushBack("jangan hapus saya")

	l.Remove(e1)

	fmt.Println("List setelah penghapusan:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("-", e.Value)
	}
}
