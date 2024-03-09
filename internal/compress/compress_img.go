package compress

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/dim-ops/mdqc/internal/mdextract"
)

func CompressImg(path string) error {
	files, err := mdextract.GetFiles(path)
	if err != nil {
		return fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}

	for _, file := range files {
		inputFile, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("impossible to open file %s: %w", file, err)
		}
		defer inputFile.Close()

		img, _, err := image.Decode(inputFile)
		if err != nil {
			return fmt.Errorf("impossible to read file %s: %w", file, err)
		}

		outputFile, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("impossible to create file %s: %w", file, err)
		}
		defer outputFile.Close()

		ext := strings.ToLower(filepath.Ext(file))
		if ext == ".jpg" || ext == ".jpeg" {
			err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: 80})
			if err != nil {
				return fmt.Errorf("impossible to compress image %s: %v", file, err)
			}
		} else if ext == ".png" {
			err = png.Encode(outputFile, img)
			if err != nil {
				return fmt.Errorf("impossible to compress image %s: %v", file, err)
			}
		} else {
			fmt.Printf("impossible to compress this image format for %s\n", file)
			continue
		}
	}
	return nil
}
