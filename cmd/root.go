package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskd",
	Short: "A distributed task scheduler",
	Long: `A distributed task scheduler that allows you to schedule and manage tasks
across multiple machines with reliable execution and monitoring.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
