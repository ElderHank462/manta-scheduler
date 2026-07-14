package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"manta/models"
	"manta/util"

	"github.com/spf13/cobra"
)

var desc string
var duration int
var due string
var group string = ""

var timeFormats []string = []string{
	"2006-01-02",
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task with a given name.",
	Long: `Create a new named task. 
	
	Additional tasks attributes can be defined using --desc for description, 
	--dur for days to complete, and --due for the task's due date (defaults 
	to one week from time of creation).`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// create the new task variable
		newTask := models.Task{
            Name:        args[0],
            Description: desc,
            Duration:    duration,
            DueDate:     due,
			Group:		group,
        }

		tasksPath := util.GetTasksPath()

        // load existing tasks
        tasks := []models.Task{}
        content, err := os.ReadFile(tasksPath)
		if err != nil {
			// tasks.json has not been created yet, so initialize it as an empty json file
			if os.IsNotExist(err) {
				content = []byte("[]")
			} else {
				fmt.Println("Error reading json file: ", err)
				return
			}

		}

		// write the tasks to the go array
        json.Unmarshal(content, &tasks)

        // append the new task to the tasks array
        tasks = append(tasks, newTask)
        updatedJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to json: ", err)
			return
		}

        if err := os.WriteFile(tasksPath, updatedJSON, 0644); err != nil {
			fmt.Println("Error writing to json file: ", err)
			return
		}

        fmt.Printf("Task '%s' added successfully!\n", args[0])


	},
}

func init() {
	rootCmd.AddCommand(addCmd)


	addCmd.Flags().StringVar(&desc, "desc", "No description", "Textual description of the task")
	addCmd.Flags().IntVar(&duration, "dur", 1, "Days required to complete task")

	var defaultTime = time.Now().AddDate(0, 0, 7)
	addCmd.Flags().StringVar(&due, "due", defaultTime.Format("2006-01-02"), "Due date")
}
