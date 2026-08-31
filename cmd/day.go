package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"manta/models"
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
			// list assigned tasks
			util.PrintBorder()
			util.PrintTitle()


			tasks := tasksForDate(targetDay)

			text := util.FormatHeader("Tasks for: " + targetDay)
			util.PrintFormattedLn(text)

			printTasks(tasks)

			util.PrintFormattedLn("Available commands: q (quit), a (assign), r (remove), s (search)")

			util.PrintBorder()
			
			scanner := bufio.NewScanner(os.Stdin)
			
			// input loop
			for {
				// display commands
				util.PrintFormattedLn("Enter a command...")
	
				// listen for input
				fmt.Print(":")
				if !scanner.Scan() {
					break
				}
	
				// process input
				input := strings.TrimSpace(scanner.Text())

				switch input {
					case "q":
						util.PrintFormattedLn("Quitting edit mode...")
						return
					case "x":
						util.PrintFormattedLn("Cancelled command")
					case "a":
						util.PrintFormattedLn("assign called")
					case "r":
						util.PrintFormattedLn("remove called")
					case "s":
						util.PrintFormattedLn("search called")
				}

			}



		} else {

			util.PrintBorder()
			util.PrintTitle()



			tasks := tasksForDate(targetDay)

			completed := filterTasksByCompletion(true, tasks)
			incomplete := filterTasksByCompletion(false, tasks)

			title := util.FormatHeader("Tasks for: " + targetDay)
			completed_header := util.FormatHeader("Already completed:")
			
			if len(tasks) == 0 {
				util.PrintFormattedLn("No tasks assigned for %s", targetDay)
			} else {
				util.PrintFormattedLn(title)
				printTasks(incomplete)
				util.PrintFormattedLn(completed_header)
				printTasks(completed)
			}


			util.PrintBorder()
		}
	},
}

func printTasks(tasks []models.Task) {
	// display date (and unique identifier, random sea creature?)
	
	for i := range tasks {
		task := tasks[i]
		combined := "%d. %s: %s (Due: %s, Days Required: %d)" 

		

		util.PrintFormattedLn(combined, i + 1, task.Name, task.Description, task.DueDate, task.Duration)
	}

}

func tasksForDate(dateString string) []models.Task {

	base := util.LoadTasks()
	var filtered []models.Task

	for _, task := range base {
		if task.Assigned == dateString {
			filtered = append(filtered, task)
		}
	}

	slices.SortFunc(filtered, func(a, b models.Task) int {
		return a.Duration - b.Duration
	})

	return filtered
}

func filterTasksByAssignment(filter string) []models.Task {
	base := util.LoadTasks()
	var filtered []models.Task

	for _, task := range base {
		if task.Assigned == filter {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func filterTasksByCompletion(filter bool, base []models.Task) []models.Task {
	if base == nil {
		base = util.LoadTasks()
	}

	var filtered []models.Task

	for _, task := range base {
		if task.Completed == filter {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

// Function to handle the assign command within edit mode
func assign(date string) {
	// filter list of tasks by assigned=""
	tasks := filterTasksByAssignment("")

	slices.SortFunc(tasks, func(a, b models.Task) int {
		dateA, err := time.Parse("2006-01-02", a.DueDate)
		if err != nil {
			return 0 // suppress
		}

		dateB, err := time.Parse("2006-01-02", b.DueDate)
		if err != nil {
			return 0 // suppress
		}

		if dateA.Before(dateB) {
			return -1
		} else if dateA.Equal(dateB) {
			return 0
		} else {
			return 1
		}

	})

	// display list (numbered)
	printTasks(tasks)


	// prompt for index
	scanner := bufio.NewScanner(os.Stdin)
	assignIndex := -1

	for {
		fmt.Print("Specify a task's number to assign it: ")

		if !scanner.Scan() {
			break
		}
	
		// process input
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "x" {
			util.PrintFormattedLn("Cancelled task assignment")
			break
		}

		n, err := strconv.Atoi(input)
		if err != nil || n <= 0 {    // <- rejects non-integers AND zero/negative values
			fmt.Println("Please enter a valid positive number")
		} else {
			assignIndex = n
			break
		}
	}

	if assignIndex == -1 {
		return
	} else {
		// assign index and return
		assigned := tasks[assignIndex-1]

		assigned.Assigned = date
		return
	}

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
