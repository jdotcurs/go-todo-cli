package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jdotcurs/go-todo-cli/internal/db"
	"github.com/jdotcurs/go-todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task",
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

		var task models.Task
		result := db.DB.First(&task, id)
		if result.Error != nil {
			fmt.Println("Task not found")
			return
		}

		description, _ := cmd.Flags().GetString("description")
		dueDateStr, _ := cmd.Flags().GetString("due")

		if description != "" {
			task.Description = description
		}

		if dueDateStr != "" {
			dueDate, err := time.Parse("2006-01-02", dueDateStr)
			if err != nil {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD")
				return
			}
			task.DueDate = dueDate
		}

		result = db.DB.Save(&task)
		if result.Error != nil {
			fmt.Println("Error updating task:", result.Error)
			return
		}

		fmt.Println("Task updated successfully!")
	},
}

func init() {
	editCmd.Flags().StringP("description", "d", "", "Task description")
	editCmd.Flags().StringP("due", "t", "", "Due date (YYYY-MM-DD)")
}
