package cmd

import (
	"fmt"
	"strconv"

	"github.com/barmstrong9/brandon/gophercises/todo/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete and removes it.",
	Run: func(cmd *cobra.Command, args []string) {
		var taskIds []int
		for _, arg := range args {
			//strconv.Atoi is string conversion, Ascii to Int
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				taskIds = append(taskIds, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something Went Wrong...", err)
			return
		}
		for _, id := range taskIds {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid Task Number", id)
				continue
			}
			task := tasks[id-1]
			err = db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed To Mark \"%d\"as completed. Error:%s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\"as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
