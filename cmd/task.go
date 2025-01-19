package cmd

import (
	"fmt"
	"task-scheduler/internal/scheduler"
	"time"

	"github.com/spf13/cobra"
)

var (
	schedule string
	timeout  time.Duration
)

func init() {
	taskCmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
		Long:  `Create, list, and manage scheduled tasks.`,
	}

	createCmd := &cobra.Command{
		Use:   "create [command]",
		Short: "Create a new task",
		Args:  cobra.ExactArgs(1),
		RunE:  createTask,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE:  listTasks,
	}

	// Add flags to create command
	createCmd.Flags().StringVarP(&schedule, "schedule", "s", "", "Cron schedule expression")
	createCmd.Flags().DurationVarP(&timeout, "timeout", "t", 30*time.Minute, "Task timeout duration")

	// Add commands to task command
	taskCmd.AddCommand(createCmd)
	taskCmd.AddCommand(listCmd)

	// Add task command to root
	rootCmd.AddCommand(taskCmd)
}

func createTask(cmd *cobra.Command, args []string) error {
	command := args[0]

	scheduler.ScheduleTask(schedule, command)

	return nil
}

func listTasks(cmd *cobra.Command, args []string) error {
	tasks := scheduler.GetTasks()

	for _, task := range tasks {
		fmt.Println(task)
	}

	return nil
}
