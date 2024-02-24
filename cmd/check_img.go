package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Checks images link(s) in your markdown file(s)",
	Long:  `Checks that the links to your images in your markdown files work.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := getFiles()
		if err != nil {
			return fmt.Errorf("impossible to get file(s): %w", err)
		}

		imgLinks, err := getLinks(files, regexImg)
		if err != nil {
			return fmt.Errorf("impossible to get image(s) link(s): %w", err)
		}

		var wg sync.WaitGroup
		results := make(chan string)

		for _, link := range imgLinks {
			wg.Add(1)
			go checkImgLinks(link, results, &wg)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		for i := 0; i < len(imgLinks); i++ {
			fmt.Print(<-results)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	imgCmd.PersistentFlags().StringVarP(&imgPath, "imgPath", "i", "", "path to image(s)")
	imgCmd.MarkPersistentFlagRequired("imgPath")
}

func checkImgLinks(link string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	cheminImage := filepath.Join(imgPath, link)
	if _, err := os.Stat(cheminImage); os.IsNotExist(err) {
		results <- fmt.Sprintf("Error checking link %s: %s\n", link, err)
	}
}
