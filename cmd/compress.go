package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/kaustubha-chaturvedi/yst-img/internal/pipeline"
	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

var maxSize string
var compressCmd = &cobra.Command{
	Use:   "compress <input> [output]>",
	Short: "Compress images (auto batch if directory)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Usage()
			return fmt.Errorf("Missing required arguments :required input")
		}
		input := args[0]
		
		var output string
		if len(args) >= 2 {
			output = args[1]
		} else {
			output = defaultOutput(input,"_compressed")
		}
		return pipeline.Run(
			input,
			output,
			quality,
			workers,
			format,
			maxSize,
			formats.Compress,
		)
	},
}

func init() {
	compressCmd.Flags().IntVarP(&quality, "quality", "q", 0, "quality (0 = auto)")
	compressCmd.Flags().IntVarP(&workers, "workers", "w", 4, "parallel workers")
	compressCmd.Flags().StringVarP(&maxSize,"max-size", "m","","target max output size (e.g. 300k, 2m)",)
	
}