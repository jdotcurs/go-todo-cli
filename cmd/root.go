package cmd

import (
	"github.com/jdotcurs/go-todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI todo application",
	Long:  `A simple CLI todo application built with Go and SQLite.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		db.InitDB()
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(editCmd)
}
