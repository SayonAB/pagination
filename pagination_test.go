package pagination

import (
	"testing"
)

func Test_CurrentPageNormalization(t *testing.T) {
	zero := New(10, 2, 0)
	if zero.CurrentPage() != 1 {
		t.Errorf("New did not normalize current page to 1 based index")
	}

	upper := New(10, 2, 7)
	if upper.CurrentPage() != 5 {
		t.Errorf("New did not normalize current page to last page on overflow")
	}

	none := New(73, 25, 2)
	if none.CurrentPage() != 2 {
		t.Errorf("New did normalize current page when not needed")
	}
}

func Test_Offset(t *testing.T) {
	p := New(28, 25, 2)
	expected := 25
	got := p.Offset()
	if got != expected {
		t.Errorf("Paginator#Offset expected %d, got %d", expected, got)
	}

	p = New(10, 3, 1)
	expected = 0
	got = p.Offset()
	if got != expected {
		t.Errorf("Paginator#Offset expected %d, got %d", expected, got)
	}

	p = New(10, 3, 4)
	expected = 9
	got = p.Offset()
	if got != expected {
		t.Errorf("Paginator#Offset expected %d, got %d", expected, got)
	}
}

func Test_NumberOfPages(t *testing.T) {
	p := New(28, 25, 2)
	expected := 2
	got := p.NumberOfPages()
	if got != expected {
		t.Errorf("Paginator#NumberOfPages expected %d, got %d", expected, got)
	}

	p = New(10, 3, 1)
	expected = 4
	got = p.NumberOfPages()
	if got != expected {
		t.Errorf("Paginator#NumberOfPages expected %d, got %d", expected, got)
	}

	p = New(10, 25, 1)
	expected = 1
	got = p.NumberOfPages()
	if got != expected {
		t.Errorf("Paginator#NumberOfPages expected %d, got %d", expected, got)
	}
}

func Test_PreviousPage(t *testing.T) {
	p := New(28, 25, 2)
	expected := 1
	got := p.PreviousPage()
	if got != expected {
		t.Errorf("Paginator#PreviousPage expected %d, got %d", expected, got)
	}

	p = New(10, 3, 1)
	expected = 1
	got = p.PreviousPage()
	if got != expected {
		t.Errorf("Paginator#PreviousPage expected %d, got %d", expected, got)
	}

	p = New(101, 25, 5)
	expected = 4
	got = p.PreviousPage()
	if got != expected {
		t.Errorf("Paginator#PreviousPage expected %d, got %d", expected, got)
	}
}

func Test_NextPage(t *testing.T) {
	p := New(28, 25, 2)
	expected := 2
	got := p.NextPage()
	if got != expected {
		t.Errorf("Paginator#NextPage expected %d, got %d", expected, got)
	}

	p = New(10, 3, 1)
	expected = 2
	got = p.NextPage()
	if got != expected {
		t.Errorf("Paginator#NextPage expected %d, got %d", expected, got)
	}

	p = New(101, 25, 5)
	expected = 5
	got = p.NextPage()
	if got != expected {
		t.Errorf("Paginator#NextPage expected %d, got %d", expected, got)
	}
}

func Test_IsCurrentPage(t *testing.T) {
	p := New(28, 25, 2)
	expected := true
	got := p.IsCurrentPage(2)
	if got != expected {
		t.Errorf("Paginator#IsCurrentPage expected %v, got %v", expected, got)
	}

	p = New(10, 3, 1)
	expected = false
	got = p.IsCurrentPage(2)
	if got != expected {
		t.Errorf("Paginator#IsCurrentPage expected %v, got %v", expected, got)
	}

	p = New(10, 3, 1)
	expected = false
	got = p.IsCurrentPage(200)
	if got != expected {
		t.Errorf("Paginator#IsCurrentPage expected %v, got %v", expected, got)
	}
}

func Test_Pages(t *testing.T) {
	p := New(28, 25, 2)
	expectedLength := 2
	got := p.Pages()
	if len(got) != expectedLength {
		t.Errorf("Paginator#Pages expected %d, got %d", expectedLength, len(got))
	}

	p = New(10, 3, 1)
	expectedLength = 4
	got = p.Pages()
	if len(got) != expectedLength {
		t.Errorf("Paginator#Pages expected %d, got %d", expectedLength, len(got))
	}

	p = New(10, 25, 1)
	expectedLength = 1
	got = p.Pages()
	if len(got) != expectedLength {
		t.Errorf("Paginator#NextPage expected %d, got %d", expectedLength, len(got))
	}
}

func Test_PagesStream(t *testing.T) {
	p := New(28, 25, 2)
	result := make([]int, 0, 5)
	expectedLength := 2
	for i := range p.PagesStream() {
		result = append(result, i)
	}
	if len(result) != expectedLength {
		t.Errorf("Paginator#Pages expected %d, got %d", expectedLength, len(result))
	}
}

func Test_Show(t *testing.T) {
	p := New(28, 25, 2)
	expected := true
	got := p.Show()
	if got != expected {
		t.Errorf("Paginator#Show expected %v, got %v", expected, got)
	}

	p = New(10, 25, 1)
	expected = false
	got = p.Show()
	if got != expected {
		t.Errorf("Paginator#Show expected %v, got %v", expected, got)
	}
}
