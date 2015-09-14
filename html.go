package pagination

import (
	"bytes"
	"fmt"
	"html/template"
)

// The template used to generate the pagination list.
const tmpl string = `
{{ if .Show }}
  <div class="pagination">
    <ul class="pagination-list">
      <li class="pagination-item-special">
        <a class="pagination-link pagination-previous" href="?page={{.PreviousPage}}">&lt;-</a>
      </li>
      {{ range $page := .PagesStream }}
        <li class="pagination-item">
          <a class="pagination-link{{if $.IsCurrentPage $page}} pagination-current {{end}}" href="?page={{$page}}">{{$page}}</a>
        </li>
      {{end}}
      <li class="pagination-item-special">
        <a class="pagination-link pagination-next" href="?page={{.NextPage}}">-&gt;</a>
      </li>
    </ul>
  </div>
{{end}}
`

// HTML is a specialised paginator that also knows how to represent the pagination as a HTML list.
type HTML struct {
	*Pagination
}

// NewHTML returns a new html pagination with the provided values.
func NewHTML(numberOfItems, itemsPerPage, currentPage int) *HTML {
	return &HTML{Pagination: New(numberOfItems, itemsPerPage, currentPage)}
}

// Render returns the HTML representation of the pagination.
func (h *HTML) Render() template.HTML {
	var out bytes.Buffer
	t := template.Must(template.New("pagination").Parse(tmpl))
	err := t.Execute(&out, h)
	if err != nil {
		return template.HTML(fmt.Sprintf("Error executing pagination template: %s", err))
	}
	return template.HTML(out.String())
}
