package main

import (
	"os"

	"github.com/wham/minimalist-go-react-app/v2/internal/ui"
)

func main() {
	data := ui.BuildForProduction()

	os.WriteFile("build/ui.js", data, 0644)
}
