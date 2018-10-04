package cmd

import (
	"fmt"
	"os"

	"github.com/barmstrong9/brandon/gophercises/todo/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the current tasks that need doing.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something Went Wrong...", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You Have No Tasks")
			return
		}
		fmt.Println("You Have The Following Tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}