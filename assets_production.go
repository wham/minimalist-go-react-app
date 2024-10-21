//go:build !development

package assets

import (
	"embed"
)

//go:embed static/*
var StaticFS embed.FS

//go:embed templates/*
var TemplatesFS embed.FS

//go:embed build/ui.js
var js []byte

func ReadJS() []byte {
	return js
}
