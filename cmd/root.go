package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// version is injected at build time via -ldflags="-X manta/cmd.version=..."
var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "manta",
	Short: "A scheduling app for tracking and planning tasks.",
	Long:  `A scheduling app for tracking and planning tasks.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
