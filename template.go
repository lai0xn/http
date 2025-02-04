package http

import (
	"os"
)

// template represents an HTML template with a file path and its content.
type template struct {
	path    string // Path to the template file
	content string // Content of the template file
}

// NewTemplate creates and returns a new template instance with the specified file path.
func NewTemplate(path string) *template {
	return &template{
		path: path,
	}
}

// Parse reads the template file from the specified path and stores its content.
// It returns an error if the file cannot be read.
func (t *template) Parse() error {
	file, err := os.ReadFile(t.path)
	if err != nil {
		return err
	}
	t.content = string(file) // Store the file content as a string
	return nil
}

// Execute writes the template content to the provided ResponseWriter.
// It sets the "Content-Type" header to "text/html" and writes the template content as the response body.
func (t *template) Execute(w ResponseWriter) {
	w.WriteHeader("Content-Type", "text/html") // Set the response content type
	w.WriteString(t.content)                   // Write the template content to the response
}
