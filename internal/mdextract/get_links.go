package mdextract

import (
	"bufio"
	"os"
	"regexp"
)

func GetLinks(files []string, regex *regexp.Regexp) (webLinks []string, err error) {

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
