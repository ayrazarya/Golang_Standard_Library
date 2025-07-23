package package_container_list

import (
	"container/list"
	"fmt"
)

func DemoReverseTraversal() {
	fmt.Println("\n=== Demo Reverse Traversal ===")

	l := list.New()
	l.PushBack("X")
	l.PushBack("Y")
	l.PushBack("Z")

	fmt.Println("Dari belakang ke depan:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println("-", e.Value)
	}
}
