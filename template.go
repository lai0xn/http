package http

import (
	"fmt"
	"os"
)

type template struct {
	path    string
	content string
}

func NewTemplate(path string) *template {
	return &template{
		path: path,
	}
}

func (t *template) Parse() error {
	file, err := os.ReadFile(t.path)
	if err != nil {
		return err
	}
	t.content = string(file)

	return nil

}

func (t *template) Execute(w ResponseWriter) {
	fmt.Println("cc")
	w.WriteHeader("Content-Type", "text/html")
	w.WriteString(t.content)
}
