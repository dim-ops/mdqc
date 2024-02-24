package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdqc",
	Short: "A simple program to check your markdown file(s)",
	Long:  "This program checks that the links to your images and web work well, other features are in development",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to target file(s)")
	rootCmd.MarkPersistentFlagRequired("path")
}
