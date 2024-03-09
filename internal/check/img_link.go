package check

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/dim-ops/mdqc/internal/mdextract"
)

func CheckImgLinks(path, imgPath string) error {
	files, err := mdextract.GetFiles(path)
	if err != nil {
		return fmt.Errorf("impossible to get file(s): %w", err)
	}

	imgLinks, err := mdextract.GetLinks(files, regexImg)
	if err != nil {
		return fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, 1)

	for _, link := range imgLinks {
		wg.Add(1)
		go checkOneImgLinks(link, imgPath, errs, &wg)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for i := 0; i < len(imgLinks); i++ {
		if err := <-errs; err != nil {
			return err
		}
	}
	return nil
}

func checkOneImgLinks(link, imgPath string, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	cheminImage := filepath.Join(imgPath, link)
	if _, err := os.Stat(cheminImage); os.IsNotExist(err) {
		errs <- fmt.Errorf("Error checking link %s: %s\n", link, err)
	}
}
