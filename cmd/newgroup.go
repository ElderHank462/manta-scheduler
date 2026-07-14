package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"manta/models"
	"manta/util"

	"github.com/spf13/cobra"
)

var groupColor string


// newgroupCmd represents the newgroup command
var newgroupCmd = &cobra.Command{
	Use:   "newgroup",
	Short: "Create a new task group.",
	Long: `Create a new named task group with an optional color. 
The color is used to display the group in task listings.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		groupsPath := util.GetGroupsPath()
		
		groups := make(map[string]models.TaskGroup)
		content, err := os.ReadFile(groupsPath)
		if err != nil {
			if !os.IsNotExist(err) {
				fmt.Println("Error reading json file: ", err)
				return
			}
		} else {
			if err := json.Unmarshal(content, &groups); err != nil {
				fmt.Println("Error unmarshaling task groups: ", err)
				return
			}
		}

		newGroup := models.TaskGroup{
			Name: args[0],
			Color: groupColor,
		}

		if _, exists := groups[args[0]]; exists {
			fmt.Printf("Group '%s' already exists.\n", args[0])
			return
		}

		groups[args[0]] = newGroup
		updatedJSON, err := json.MarshalIndent(groups, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to json: ", err)
			return
		}

		if err := os.WriteFile(groupsPath, updatedJSON, 0644); err != nil {
			fmt.Println("Error writing to json file: ", err)
			return
		}

		fmt.Printf("Task group '%s' successfully created!\n", args[0])

	},
}

func init() {
	rootCmd.AddCommand(newgroupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newgroupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newgroupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	
	newgroupCmd.Flags().StringVar(&groupColor, "color", "#6038BC", "The color for the group to display with.")
}
