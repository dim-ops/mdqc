package cmd

import (
	"fmt"

	"github.com/dim-ops/mdqc/internal/compress"
	"github.com/spf13/cobra"
)

var compressImg = &cobra.Command{
	Use:   "compress",
	Short: "Compress your img(s)",
	Long:  `Compress your img(s) to improve their size and SEO`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := compress.CompressImg(path)
		if err != nil {
			return fmt.Errorf("Impossible to crompress image(s) in %s: %w", path, err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressImg)
}
