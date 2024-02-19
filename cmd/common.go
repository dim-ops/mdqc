package cmd

import (
	"bufio"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

func getLinks(files []fs.DirEntry, regex *regexp.Regexp) ([]string, error) {

	var webLinks []string

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".md") {
			f, err := os.Open(path + "/" + file.Name())
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
	}
	return webLinks, nil
}
