package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd allows the creation of commands for the CLI
var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a CLI task manager",
}
