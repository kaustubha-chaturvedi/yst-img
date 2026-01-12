package cmd

import (
	"fmt"

	"github.com/kaustubha-chaturvedi/yst-img/internal/pipeline"
	"github.com/spf13/cobra"
)

var width int

var resizeCmd = &cobra.Command{
	Use:   "resize <input> [output]",
	Short: "Resize an image",
	RunE: func(cmd *cobra.Command, args []string) error {
		if width <= 0 {
			return fmt.Errorf("invalid width")
		}
		if len(args) < 1 {
			cmd.Usage()
			return fmt.Errorf("Missing required arguments :required input")
		}
		input := args[0]

		var output string
		if len(args) >= 2 {
			output = args[1]
		} else {
			output = defaultOutput(input,"_resized")
		}
		return pipeline.Resize(input, output, width)
	},
}

func init() {
	resizeCmd.Flags().IntVarP(&width, "width", "w", 0, "width of image")
}