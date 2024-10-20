//go:build !development

package assets

import (
	"text/template"

	embedx "github.com/wham/minimalist-go-react-app/v2"
)

func ReadStaticFile(name string) ([]byte, error) {
	return embedx.Static.ReadFile("static/" + name)
}

func ParseTemplate(name string) (*template.Template, error) {
	return template.ParseFS(embedx.Templates, "templates/"+name)
}
