package pagination_test

import (
	"fmt"

	"github.com/SayonAB/pagination"
)

func ExamplePagination_PagesStream() {
	p := pagination.New(110, 25, 1)
	pages := make([]int, 0, 5)
	for i := range p.PagesStream() {
		pages = append(pages, i)
	}
	fmt.Printf("%v", pages)
	// Output:
	// [1 2 3 4 5]
}
