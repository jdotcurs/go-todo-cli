package cmd

import (
	"fmt"
	"time"

	"github.com/jdotcurs/go-todo-cli/internal/db"
	"github.com/jdotcurs/go-todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		dueDateStr, _ := cmd.Flags().GetString("due")

		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("Invalid date format. Please use YYYY-MM-DD")
			return
		}

		task := models.Task{
			Description: description,
			DueDate:     dueDate,
		}

		result := db.DB.Create(&task)
		if result.Error != nil {
			fmt.Println("Error creating task:", result.Error)
			return
		}

		fmt.Println("Task created successfully!")
	},
}

func init() {
	createCmd.Flags().StringP("description", "d", "", "Task description")
	createCmd.Flags().StringP("due", "t", "", "Due date (YYYY-MM-DD)")
	createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagRequired("due")
}
