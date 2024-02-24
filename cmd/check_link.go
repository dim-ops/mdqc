package cmd

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Checks web link(s) in your markdown file(s)",
	Long:  `Checks that the links to web sites in your markdown files work.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := getFiles()
		if err != nil {
			return fmt.Errorf("impossible to get image(s) link(s): %w", err)
		}

		webLinks, err := getLinks(files, regexLink)
		if err != nil {
			return fmt.Errorf("impossible to get web link(s): %w", err)
		}

		fmt.Println(webLinks)
		var wg sync.WaitGroup
		results := make(chan string)

		for _, link := range webLinks {
			wg.Add(1)
			fmt.Println(link)
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
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}

func checkWebLinks(link string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		results <- fmt.Sprintf("Error checking link %s: %s\n", link, err)
	}
}
