package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "A distributed task scheduler",
	Long: `A distributed task scheduler that allows you to schedule and manage tasks
across multiple machines with reliable execution and monitoring.`,
}

var (
	schedule string
	timeout  time.Duration
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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
		//		RunE:  createTask,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		// 		RunE:  listTasks,
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
