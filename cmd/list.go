package cmd

import (
	"fmt"
	"time"

	"github.com/jdotcurs/go-todo-cli/internal/db"
	"github.com/jdotcurs/go-todo-cli/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []models.Task
		result := db.DB.Find(&tasks)
		if result.Error != nil {
			fmt.Println("Error fetching tasks:", result.Error)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %d\nDescription: %s\nCreated At: %s\nDue Date: %s\n\n",
				task.ID, task.Description, task.CreatedAt.Format(time.RFC822), task.DueDate.Format("2006-01-02"))
		}
	},
}
