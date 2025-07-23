package package_sort

import (
	"fmt"
	"sort"
)

func CustomSort() {
	type Product struct {
		Name  string
		Price int
	}

	products := []Product{
		{"Mouse", 100},
		{"Keyboard", 300},
		{"Monitor", 200},
	}

	// Sort berdasarkan harga tertinggi ke terendah
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})

	fmt.Println("Sorted by descending price:")
	for _, p := range products {
		fmt.Printf("%s: %d\n", p.Name, p.Price)
	}
}
