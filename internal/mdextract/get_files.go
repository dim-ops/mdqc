package mdextract

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func GetFiles(path string) (files []string, err error) {
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if isValidExtension(path) {
			files = append(files, path)
		}
		if err != nil {
			return fmt.Errorf("impossible to walk dir: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("impossible to get file(s): %w", err)
	}
	return files, nil
}

func isValidExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".md"
}
