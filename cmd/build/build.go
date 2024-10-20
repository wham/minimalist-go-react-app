package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"ui/main.tsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
	})

	os.WriteFile("build/ui.js", result.OutputFiles[0].Contents, 0644)
}
