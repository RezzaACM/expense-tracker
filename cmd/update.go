package cmd

import (
	"expanse-tracker/internal/models"
	"expanse-tracker/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewJSONStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return err
		}

		var newExpenses []models.Expense
		found := false
		for _, expense := range expenses {
			if expense.ID == id {
				found = true
				expense.Name = name
				expense.Description = description
				expense.Amount = amount
				newExpenses = append(newExpenses, expense)
				continue
			}
			newExpenses = append(newExpenses, expense)
		}

		if !found {
			return fmt.Errorf("expense with id %s not found", id)
		}

		err = store.SaveExpenses(newExpenses)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	updateCmd.Flags().StringVar(&id, "id", "", "ID of the expense to update")
	updateCmd.Flags().StringVar(&name, "name", "", "Name of the expense")
	updateCmd.Flags().StringVar(&description, "description", "", "Description of the expense")
	updateCmd.Flags().Float64Var(&amount, "amount", 0, "Amount of the expense")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)
}
