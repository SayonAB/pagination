package pagination

import (
	"bytes"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func assertHasHtmlElements(t *testing.T, html *goquery.Document, selector string, num int) {
	got := html.Find(selector)
	if got.Length() != num {
		t.Errorf("Expected to find %d html elements (%s) but found %d", num, selector, got.Length())
	}
}

func Test_Render(t *testing.T) {
	p := NewHTML(28, 25, 5)
	output := string(p.Render())
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(output))
	if err != nil {
		t.Fatal(err)
	}

	assertHasHtmlElements(t, doc, ".pagination-list", 1)
	assertHasHtmlElements(t, doc, ".pagination-item-special", 2)
	assertHasHtmlElements(t, doc, ".pagination-link", 2+2) // Previous and next + 2 pages (1 and 2)
	assertHasHtmlElements(t, doc, ".pagination-current", 1)
}

func Test_RenderNothing(t *testing.T) {
	p := NewHTML(1, 25, 1)
	output := string(p.Render())
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(output))
	if err != nil {
		t.Fatal(err)
	}

	assertHasHtmlElements(t, doc, ".pagination-list", 0)
	assertHasHtmlElements(t, doc, ".pagination-item-special", 0)
	assertHasHtmlElements(t, doc, ".pagination-link", 0)
	assertHasHtmlElements(t, doc, ".pagination-current", 0)
}
