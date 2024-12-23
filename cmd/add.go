package cmd

import (
	"expanse-tracker/internal/models"
	"expanse-tracker/internal/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	name        string
	description string
	amount      float64
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewJSONStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return err
		}

		newExpense := models.Expense{
			ID:          uuid.New().String(),
			Name:        name,
			Description: description,
			Date:        time.Now(),
			Amount:      amount,
		}

		expenses = append(expenses, newExpense)
		if err := store.SaveExpenses(expenses); err != nil {
			return err
		}

		fmt.Printf("Expense added with ID %s\n", newExpense.ID)
		return nil

	},
}

func init() {
	addCmd.Flags().StringVar(&name, "name", "", "Name of the expense")
	addCmd.Flags().StringVar(&description, "description", "", "Description of the expense")
	addCmd.Flags().Float64Var(&amount, "amount", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("amount")
	rootCmd.AddCommand(addCmd)
}
