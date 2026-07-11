package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"manta/models"
	"manta/util"

	"github.com/spf13/cobra"
)

// assignCmd represents the assign command
var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a task to a group",
	Long: `Assign a task to a group. The first argument is the task name, 
the second argument is the group name.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		groupName := args[1]

		tasksPath := util.GetTasksPath()

		tasks := []models.Task{}
		content, err := os.ReadFile(tasksPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No tasks file found.")
				return
			}
			fmt.Println("Error reading json file: ", err)
			return
		}

		if err := json.Unmarshal(content, &tasks); err != nil {
			fmt.Println("Error reading tasks: ", err)
			return
		}

		found := false
		for index, t := range tasks {
			if t.Name == taskName {
				tasks[index].Group = groupName
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Task '%s' not found.\n", taskName)
			return
		}

		updatedJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to json: ", err)
			return
		}

		if err := os.WriteFile(tasksPath, updatedJSON, 0644); err != nil {
			fmt.Println("Error writing to json file: ", err)
			return
		}

		fmt.Printf("Task '%s' assigned to group '%s'!\n", taskName, groupName)
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
