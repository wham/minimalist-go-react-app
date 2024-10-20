//go:build development

package assets

import (
	"os"
	"text/template"
)

func ReadStaticFile(name string) ([]byte, error) {
	return os.ReadFile("static/" + name)
}

func ParseTemplate(name string) (*template.Template, error) {
	fs := os.DirFS("templates")
	return template.ParseFS(fs, name)
}
