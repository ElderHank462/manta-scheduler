package cmd

import (
	"fmt"
	"time"

	"manta/util"

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
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var targetDay string

		if len(args) != 0 {
			targetDay = args[0]
		} else {
			targetDay = time.Now().Format("2006-01-02")
		}
		
		if editMode {
			fmt.Println("entering edit mode...")
		} else {

			// display date (and unique identifier, random sea creature?)
			text := util.FormatHeader("Tasks for: " + targetDay)

			util.PrintFormattedLn(text)

			// display tasks, ordered by duration
			util.PrintFormattedLn("Tasks go here :)")
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
