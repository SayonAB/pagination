[![Build Status](https://img.shields.io/travis/SayonAB/pagination.svg?style=flat)](https://travis-ci.org/SayonAB/pagination) [![Coverage](https://img.shields.io/codecov/c/github/SayonAB/pagination.svg?style=flat)](https://codecov.io/github/SayonAB/pagination) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/SayonAB/pagination) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/SayonAB/pagination/master/LICENSE)

# Pagination

Pagination is a simple package that provides pagination. Has HTML support using the `html/template` package.

## Installation

Install using `go get`.

```shell
$ go get github.com/SayonAB/pagination
```

Then import it in your project.

```go
import "github.com/SayonAB/pagination"
```

## Usage

Create a basic pagination and check the offset for current page

```go
numPosts := db.Posts().Count()
postsPerPage := 25
currentPage := request.Query().Int("page")
p := pagination.New(numPosts, postsPerPage, currentPage)
fmt.Printf("The current offset is: %d\n", p.Offset()) // The current offset is: 75
```

Create a HTML pagination and use it in a `html/template`

```go
p := pagination.NewHTML(110, 25, 2)
data := map[string]interface{}{
  "Pager": p,
}
var out bytes.Buffer
t := template.Must(template.New("pagination-test").Parse("{{ .Pager.Render }}"))
t.Execute(&out, data)
fmt.Printf("HTML list: %s\n", out.String())
```

## Testing

Run the tests using `go test`.

```shell
$ go test
```

## Contributing

All contributions are welcome! See [CONTRIBUTING](CONTRIBUTING.md) for more info.

## License

Licensed under MIT license. See [LICENSE](LICENSE) for more information.