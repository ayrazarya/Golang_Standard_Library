package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoPushFrontBack() {
	fmt.Println("\n=== Demo PushFront & PushBack ===")

	l := list.New()
	l.PushFront("awal")
	l.PushBack("akhir")

	fmt.Println("List setelah PushFront dan PushBack:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("-", e.Value)
	}
}
