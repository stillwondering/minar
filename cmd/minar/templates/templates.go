package templates

import (
	"html/template"
	"io"
	"path/filepath"
)

func Index(w io.Writer) error {
	return parse("index.html").Execute(w, nil)
}

func parse(file string) *template.Template {
	files := []string{
		filepath.Join("templates", "base.html"),
		filepath.Join("templates", file),
	}

	return template.Must(template.New("base.html").ParseFiles(files...))
}
