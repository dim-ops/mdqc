package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmk",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&imgPath, "imgPath", "i", "", "path to image(s)")
	rootCmd.MarkPersistentFlagRequired("imgPath")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to target file(s)")
	rootCmd.MarkPersistentFlagRequired("path")
}
