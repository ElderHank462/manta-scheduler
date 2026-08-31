package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"manta/models"
	"manta/util"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [query]",
	Short: "Search and mark a task as completed.",
	Long: `List uncompleted tasks matching the query, then prompt you to enter the number of the task to mark as completed.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(args[0])

		tasks := util.LoadTasks()
		if tasks == nil {
			return
		}

		tasksPath := util.GetTasksPath()

		var results []models.Task
		for _, t := range tasks {
			if t.Completed || strings.ToLower(t.Name) == "" {
				continue
			}
			nameMatch := strings.Contains(strings.ToLower(t.Name), query)
			descMatch := strings.Contains(strings.ToLower(t.Description), query)
			if nameMatch || descMatch {
				results = append(results, t)
			}
		}

		if len(results) == 0 {
			fmt.Println("No uncompleted tasks found matching '", args[0], "'")
			return
		}

		for i, t := range results {
			fmt.Printf("%d. %s: %s\n", i+1, t.Name, t.Description)
		}

		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Which task would you like to complete?\n> ")
			if !scanner.Scan() {
				fmt.Println("No input received.")
				return
			}
			input := strings.TrimSpace(scanner.Text())
			n, err := strconv.Atoi(input)
			if err != nil || n < 1 || n > len(results) {
				util.PrintFormattedLn("Input invalid (NaN, <= 0, or out of bounds).")
				continue
			}

			target := results[n-1]

			for index, t := range tasks {
				if t == target {
					tasks[index].Completed = true
					break
				}
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

			fmt.Printf("Task '%s' completed!\n", target.Name)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
