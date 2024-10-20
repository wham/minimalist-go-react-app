//go:build development

package assets

import (
	"os"
)

func ReadStaticFile(name string) ([]byte, error) {
	return os.ReadFile("static/" + name)
}
