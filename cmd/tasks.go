package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"manta/models"
	"manta/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Aliases: []string{"t"},
	Short: "List all uncompleted tasks.",
	Long: `List all uncompleted tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasksPath := util.GetTasksPath()
		groupsPath := util.GetGroupsPath()

		// read tasks file
        tasksContent, err := os.ReadFile(tasksPath)
        if err != nil {
            fmt.Println("No tasks found.")
            return
        }

		// read groups file
		// var groups []models.TaskGroup
		var groups map[string]models.TaskGroup
		
        groupsContent, err := os.ReadFile(groupsPath)
        if groupsContent != nil {
			if err != nil {
				fmt.Println("Error reading task groups file: ", err)
				return
			}

			// parse groups
			if err := json.Unmarshal(groupsContent, &groups); err != nil {
				fmt.Println("Error reading task groups: ", err)
				return
			}
		}

        // parse tasks
        var tasks []models.Task
        if err := json.Unmarshal(tasksContent, &tasks); err != nil {
            fmt.Println("Error reading tasks: ", err)
            return
        }


        // print tasks
        fmt.Println("\n### All Tasks ###")
        for index, t := range tasks {
			groupName := "Ungrouped"	
			groupColor := "#ff0000"

			value, exists := groups[t.Group]
			
			if(exists == true) {
				myGroup := value
				
				groupName = myGroup.Name
				groupColor = myGroup.Color
			}

			r, g, b, err := util.HexToRGB(groupColor)
			if err != nil {
				fmt.Println("Error processing group color: ", err)
			}

			colorPrint := color.RGB(r, g, b).PrintfFunc()

            fmt.Printf("%d. ", index + 1)
			
			colorPrint("(%s) ", groupName)
			
			fmt.Printf("%s: %s (Due: %s, Assigned: %s, Days Required: %d)\n", 
                t.Name, t.Description, t.DueDate, t.Assigned, t.Duration)
        }
		fmt.Println()
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
