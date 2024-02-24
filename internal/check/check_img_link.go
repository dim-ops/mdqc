package check

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/dim-ops/mdqc/internal/get"
)

func CheckImgLinks(path, imgPath string) error {
	files, err := get.GetFiles(path)
	if err != nil {
		return fmt.Errorf("impossible to get file(s): %w", err)
	}

	imgLinks, err := get.GetLinks(files, regexImg)
	if err != nil {
		return fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}

	var wg sync.WaitGroup
	results := make(chan string)

	for _, link := range imgLinks {
		wg.Add(1)
		go checkImgLinks(link, imgPath, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < len(imgLinks); i++ {
		fmt.Print(<-results)
	}
	return nil
}

func checkImgLinks(link, imgPath string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	cheminImage := filepath.Join(imgPath, link)
	if _, err := os.Stat(cheminImage); os.IsNotExist(err) {
		results <- fmt.Sprintf("Error checking link %s: %s\n", link, err)
	}
}
