//go:build development

package assets

import (
	"os"
	"text/template"

	"github.com/evanw/esbuild/pkg/api"
)

func ReadStaticFile(name string) ([]byte, error) {
	return os.ReadFile("static/" + name)
}

func ParseTemplate(name string) (*template.Template, error) {
	fs := os.DirFS("templates")
	return template.ParseFS(fs, name)
}

func ReadUiJs() []byte {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"ui/main.tsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
		Sourcemap:   api.SourceMapInline,
	})

	return result.OutputFiles[0].Contents
}
