/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"manta/models"
	"manta/util"

	"github.com/spf13/cobra"
	// "github.com/fatih/color"
)

var color string

// var colorPalette = [10]{}


// newgroupCmd represents the newgroup command
var newgroupCmd = &cobra.Command{
	Use:   "newgroup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// create the new taskgroup variable
		newGroup := models.TaskGroup{
			Name: args[0],
			Color: color,
		}
		
		// get the path to the json file
		groupsPath := util.GetGroupsPath()

		// load the task groups from the file
		groups := []models.TaskGroup{}
		content, err := os.ReadFile(groupsPath)
		if err != nil {
			if os.IsNotExist(err) {
				// check if file exists yet, if not create it as an empty json file
				content = []byte("[]")
			} else {
				fmt.Println("Error reading json file: ", err)
				return
			}
		}

		json.Unmarshal(content, &groups)

		// append new task group to the array
		groups = append(groups, newGroup)
		updatedJSON, err := json.MarshalIndent(groups, "", " ")
		if err != nil {
			fmt.Println("Error marshaling to json: ", err)
			return
		}

		// write the updated array to the file
		if err := os.WriteFile(groupsPath, updatedJSON, 0644); err != nil {
			fmt.Println("Error writing to json file: ", err)
			return
		}

		// print output message
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

	
	newgroupCmd.Flags().StringVar(&color, "color", "#6038BC", "The color for the group to display with.")
}
