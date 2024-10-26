package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/wham/minimalist-go-react-app/v2/internal/ui"
)

func main() {
	slog.Info("Building UI for production...")
	data := ui.BuildForProduction()

	outputDir := "build"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		slog.Error("Failed to create directory", "error", err)
		return
	}

	outputFile := filepath.Join(outputDir, "ui.js")
	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		slog.Error("Failed to write file", "error", err)
		return
	}
	slog.Info("Wrote file:", "file", outputFile)
}
