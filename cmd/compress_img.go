package cmd

import (
	"github.com/dim-ops/mdqc/internal/compress"
	"github.com/spf13/cobra"
)

var compressImg = &cobra.Command{
	Use:           "compress",
	Short:         "Compress your img(s)",
	Long:          `Compress your img(s) to improve their size and SEO`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := compress.CompressImg(path)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressImg)
}
