package ui

import "github.com/evanw/esbuild/pkg/api"

func BuildForDevelopment() []byte {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"ui/main.tsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
		Sourcemap:   api.SourceMapInline,
	})

	return result.OutputFiles[0].Contents
}

func BuildForProduction() []byte {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"ui/main.tsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
	})

	return result.OutputFiles[0].Contents
}
