package check

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dim-ops/mdqc/internal/mdextract"
)

func CheckWebLinks(path string) error {
	files, err := mdextract.GetFiles(path)
	if err != nil {
		return fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}

	webLinks, err := mdextract.GetLinks(files, regexLink)
	if err != nil {
		return fmt.Errorf("impossible to get web link(s): %w", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, 1)

	for _, link := range webLinks {
		wg.Add(1)
		go checkOneWebLinks(link, errs, &wg)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for i := 0; i < len(webLinks); i++ {
		if err := <-errs; err != nil {
			return err
		}
	}
	return nil
}

func checkOneWebLinks(link string, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		errs <- fmt.Errorf("Error checking link %s: %s\n", link, err)
	}
}
