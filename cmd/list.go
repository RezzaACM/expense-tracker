package cmd

import (
	"expanse-tracker/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewJSONStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return err
		}

		fmt.Printf("ID\tName\tAmount\tDate\tDescription\n")
		for _, expense := range expenses {
			fmt.Printf("%s\t%s\t$%.2f\t%s\t%s\n",
				expense.ID,
				expense.Name,
				expense.Amount,
				expense.Date,
				expense.Description,
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
