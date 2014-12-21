package pagination

import (
	"math"
)

// Pagination is a general purpose pagination type, it knows how to calculate
// offset and number of pages. It also contains some utility functions
// that helps common tasks. One special utility is the PagesStream method
// that returns a channel to range over for presenting a list of all pages
// without adding them all to a slice.
type Pagination struct {
	itemsPerPage  int
	numberOfItems int
	currentPage   int
}

// New returns a new Pagination with the provided values.
// The current page is normalized to be inside the bounds
// of the available pages. So if the current page supplied
// is less than 1 the current page is normalized as 1, and if
// it is larger than the number of pages needed its normalized
// as the last available page.
func New(numberOfItems, itemsPerPage, currentPage int) *Pagination {
	if currentPage == 0 {
		currentPage = 1
	}

	n := int(math.Ceil(float64(numberOfItems) / float64(itemsPerPage)))
	if currentPage > n {
		currentPage = n
	}

	return &Pagination{
		itemsPerPage:  itemsPerPage,
		numberOfItems: numberOfItems,
		currentPage:   currentPage,
	}
}

// PagesStream returns a channel that will be incremented to
// the available number of pages. Useful to range over when
// building a list of pages.
func (p *Pagination) PagesStream() chan int {
	stream := make(chan int)
	go func() {
		for i := 1; i <= p.NumberOfPages(); i++ {
			stream <- i
		}
		close(stream)
	}()
	return stream
}

// Offset calculates the offset into the collection the current page represents.
func (p *Pagination) Offset() int {
	return (p.CurrentPage() - 1) * p.ItemsPerPage()
}

// NumberOfPages calculates the number of pages needed
// based on number of items and items per page.
func (p *Pagination) NumberOfPages() int {
	return int(math.Ceil(float64(p.NumberOfItems()) / float64(p.ItemsPerPage())))
}

// PreviousPage returns the page number of the page before current page.
// If current page is the first in the list of pages, 1 is returned.
func (p *Pagination) PreviousPage() int {
	if p.CurrentPage() <= 1 {
		return 1
	}

	return p.CurrentPage() - 1
}

// NextPage returns the page number of the page after current page.
// If current page is the last in the list of pages, the last page number is returned.
func (p *Pagination) NextPage() int {
	if p.CurrentPage() >= p.NumberOfPages() {
		return p.NumberOfPages()
	}

	return p.CurrentPage() + 1
}

// IsCurrentPage checks a number to see if it matches the current page.
func (p *Pagination) IsCurrentPage(page int) bool {
	return p.CurrentPage() == page
}

// Pages returns a list with all page numbers.
// Eg. [1 2 3 4 5]
func (p *Pagination) Pages() []int {
	s := make([]int, 0, p.NumberOfPages())

	for i := 1; i <= p.NumberOfPages(); i++ {
		s = append(s, i)
	}

	return s
}

// Show returns true if the pagination should be used.
// Ie. if there is more than one page.
func (p *Pagination) Show() bool {
	return p.NumberOfPages() > 1
}

// CurrentPage returns the current page.
func (p *Pagination) CurrentPage() int {
	return p.currentPage
}

// NumberOfItems returns the number of items.
func (p *Pagination) NumberOfItems() int {
	return p.numberOfItems
}

// ItemsPerPage returns the number of items to show per page.
func (p *Pagination) ItemsPerPage() int {
	return p.itemsPerPage
}
