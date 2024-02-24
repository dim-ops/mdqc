package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getFiles() (files []string, err error) {
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if isValidExtension(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}
	return files, nil
}

func getLinks(files []string, regex *regexp.Regexp) (webLinks []string, err error) {

	fmt.Println(files)
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ligne := scanner.Text()

			resultats := regex.FindAllStringSubmatch(ligne, -1)
			for _, resultat := range resultats {
				link := resultat[2]
				webLinks = append(webLinks, link)
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return webLinks, nil
}

func isValidExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".md"
}
