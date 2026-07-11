package cmd

import (
	"github.com/spf13/cobra"
)

// assignCmd represents the assign command
var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// parse args
		

		// load json data into a variable

		// modify the group parameter

		// write back to json


	},
}

func init() {
	rootCmd.AddCommand(assignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// assignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// assignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
