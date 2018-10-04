package cmd

import (
	"fmt"
	"strings"

	"github.com/barmstrong9/todo/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to your list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something Went Wrong...", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your todo list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
