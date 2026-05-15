/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"manta/models"

	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Aliases: []string{"t"},
	Short: "List all uncompleted tasks.",
	Long: `List all uncompleted tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("tasks called")

		// read tasks file
        content, err := os.ReadFile("tasks.json")
        if err != nil {
            fmt.Println("No tasks found.")
            return
        }

        // parse tasks
        var tasks []models.Task
        if err := json.Unmarshal(content, &tasks); err != nil {
            fmt.Println("Error reading tasks: ", err)
            return
        }

        // print tasks
        fmt.Println("--- CURRENT TASKS ---")
        for index, t := range tasks {
            fmt.Printf("%d. [%d min] %s: %s (Due: %s)\n", 
                index + 1, t.Duration, t.Name, t.Description, t.DueDate.Format("2006-01-02"))
        }
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
