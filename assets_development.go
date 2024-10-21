//go:build development

package assets

import (
	"io/fs"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

var StaticFS fs.FS
var TemplatesFS fs.FS

func init() {
	StaticFS = os.DirFS(".")
	TemplatesFS = StaticFS
}

func ReadJS() []byte {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"ui/main.tsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
		Sourcemap:   api.SourceMapInline,
	})

	return result.OutputFiles[0].Contents
}
