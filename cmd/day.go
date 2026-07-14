package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editMode bool

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if editMode {
			fmt.Println("entering edit mode...")
		} else {
			fmt.Println("no edit mode")
		}
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dayCmd.Flags().BoolVarP(&editMode, "edit", "e", false, "Enter edit mode for the day's tasks.")
}
