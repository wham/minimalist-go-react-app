package main

import (
	"log/slog"
	"os"

	"github.com/wham/minimalist-go-react-app/v2/internal/ui"
)

func main() {
	slog.Info("Building UI for production...")
	data := ui.BuildForProduction()

	os.WriteFile("build/ui.js", data, 0644)
	slog.Info("Wrote file: build/ui.js")
}
