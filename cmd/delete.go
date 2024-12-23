package cmd

import (
	"expanse-tracker/internal/models"
	"expanse-tracker/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var id string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
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
				continue
			}
			newExpenses = append(newExpenses, expense)
		}

		if !found {
			return fmt.Errorf("expense with id %s not found", id)
		}

		if err := store.SaveExpenses(newExpenses); err != nil {
			return err
		}

		fmt.Printf("Expense with ID %s deleted\n", id)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringVar(&id, "id", "", "ID of the expense to delete")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}
