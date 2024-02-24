package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var compressImg = &cobra.Command{
	Use:   "compress",
	Short: "Compress your img(s)",
	RunE: func(cmd *cobra.Command, args []string) error {

		var files []string
		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if isValidExtension(path) {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			log.Fatalf("impossible to walk directories: %s", err)
		}

		for _, file := range files {
			inputFile, err := os.Open(file)
			if err != nil {
				fmt.Printf("impossible to open file %s: %e", file, err)
			}
			defer inputFile.Close()

			fmt.Println(file)
			img, _, err := image.Decode(inputFile)
			if err != nil {
				fmt.Printf("impossible to read file %s: %e\n", file, err)
			}

			// Créer le fichier de sortie
			outputFile, err := os.Create(file)
			if err != nil {
				fmt.Printf("impossible to create file %s: %e\n", file, err)
			}
			defer outputFile.Close()

			ext := strings.ToLower(filepath.Ext(file))
			if ext == ".jpg" || ext == ".jpeg" {
				err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: 80})
				if err != nil {
					fmt.Printf("impossible to compress image %s: %e\n", file, err)
				}
			} else if ext == ".png" {
				err = png.Encode(outputFile, img)
				if err != nil {
					fmt.Printf("impossible to compress image %s: %e\n", file, err)
				}
			} else {
				fmt.Println("impossible to compress this image format")
				continue
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressImg)
}

func isValidExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}
