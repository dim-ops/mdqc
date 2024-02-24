package cmd

import (
	"fmt"

	"github.com/dim-ops/mdqc/internal/check"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Checks web link(s) in your markdown file(s)",
	Long:  `Checks that the links to web sites in your markdown files work.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := check.CheckWebLinks(path)
		if err != nil {
			return fmt.Errorf("Impossible to check web link(s) in %s: %w", path, err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
