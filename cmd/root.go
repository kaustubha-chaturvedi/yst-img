package cmd


import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:           "yst img",
	Long:          "Plugin for image operations",
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {
	return rootCmd.Execute()
}
