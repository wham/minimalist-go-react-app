//go:build !development

package assets

import "github.com/wham/minimalist-go-react-app/v2/static"

func ReadStaticFile(name string) ([]byte, error) {
	return static.Static.ReadFile(name)
}
