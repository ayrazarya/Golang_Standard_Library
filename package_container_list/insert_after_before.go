package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoInsertAfterBefore() {
	fmt.Println("\n=== Demo InsertAfter & InsertBefore ===")

	l := list.New()
	e1 := l.PushBack("satu")
	l.InsertAfter("dua", e1)
	l.InsertBefore("nol", e1)

	fmt.Println("Hasil:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("-", e.Value)
	}
}
