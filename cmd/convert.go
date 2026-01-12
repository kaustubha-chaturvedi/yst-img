package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/kaustubha-chaturvedi/yst-img/internal/pipeline"
	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

var (
	quality int
	workers int
	format  string
)

var convertCmd = &cobra.Command{
	Use:   "convert <input> <output>",
	Short: "Convert images (auto batch if directory)",
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
			output = defaultOutput(input,"_converted")
		}
		return pipeline.Run(
			input,
			output,
			quality,
			workers,
			format,
			maxSize,
			formats.Convert,
		)
	},
}

func init() {
	convertCmd.Flags().IntVarP(&quality, "quality", "q", 0, "quality (0 = auto)")
	convertCmd.Flags().IntVarP(&workers, "workers", "w", 4, "parallel workers")
	convertCmd.Flags().StringVar(&format, "format", "", "output format (batch only)")
}
