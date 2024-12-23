package cmd

import (
	"expanse-tracker/internal/storage"
	"expanse-tracker/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var month int

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of expenses",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewJSONStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return err
		}

		var total float64
		for _, expense := range expenses {
			if int(expense.Date.Month()) == month || month == 0 {
				total += expense.Amount
			}
		}
		if month == 0 {
			fmt.Printf("Total expenses: %.2f\n", total)
		} else {
			fmt.Printf("Total expenses in %s: %.2f\n", utils.MapMonth[fmt.Sprint(month)], total)
		}
		return nil
	},
}

func init() {
	summaryCmd.Flags().IntVar(&month, "month", 0, "Month to show summary for")
	rootCmd.AddCommand(summaryCmd)
}
