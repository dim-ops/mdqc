package cmd

import (
	"github.com/dim-ops/mdqc/internal/check"
	"github.com/spf13/cobra"
)

var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Checks images link(s) in your markdown file(s)",
	Long:  `Checks that the links to your images in your markdown files work.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := check.CheckImgLinks(path, imgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	imgCmd.PersistentFlags().StringVarP(&imgPath, "imgPath", "i", "", "path to image(s)")
	imgCmd.MarkPersistentFlagRequired("imgPath")
}
