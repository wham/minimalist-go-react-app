package main

import (
	"log/slog"
	"os"
	"path"

	"github.com/wham/minimalist-go-react-app/v2/internal/ui"
)

func main() {
	if err := build(); err != nil {
		os.Exit(1)
	}
}

func build() error {
	slog.Info("Building UI for production...")

	data := ui.BuildForProduction()

	outputDirectory := "build"
	if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
		slog.Error("Failed to create output directory", "error", err)
		return err
	}

	outputFile := path.Join(outputDirectory, "ui.js")
	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		slog.Error("Failed to write output file", "error", err)
		return err
	}

	slog.Info("UI built successfully", "outputFile", outputFile)

	return nil
}
