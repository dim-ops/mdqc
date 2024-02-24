package check

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dim-ops/mdqc/internal/get"
)

func CheckWebLinks(path string) error {
	files, err := get.GetFiles(path)
	if err != nil {
		return fmt.Errorf("impossible to get image(s) link(s): %w", err)
	}

	webLinks, err := get.GetLinks(files, regexLink)
	if err != nil {
		return fmt.Errorf("impossible to get web link(s): %w", err)
	}

	var wg sync.WaitGroup
	results := make(chan string)

	for _, link := range webLinks {
		wg.Add(1)
		go checkWebLinks(link, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < len(webLinks); i++ {
		fmt.Print(<-results)
	}
	return nil
}

func checkWebLinks(link string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		results <- fmt.Sprintf("Error checking link %s: %s\n", link, err)
	}
}
