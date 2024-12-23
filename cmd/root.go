package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense Tracker CLI",
}

func Execute() error {
	return rootCmd.Execute()
}
