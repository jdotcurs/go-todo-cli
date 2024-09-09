package cmd

import (
	"fmt"
	"strconv"

	"github.com/jdotcurs/go-todo-cli/internal/db"
	"github.com/jdotcurs/go-todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a task ID")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		result := db.DB.Delete(&models.Task{}, id)
		if result.Error != nil {
			fmt.Println("Error deleting task:", result.Error)
			return
		}

		if result.RowsAffected == 0 {
			fmt.Println("Task not found")
			return
		}

		fmt.Println("Task deleted successfully!")
	},
}
