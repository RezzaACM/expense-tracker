package models

import "time"

type Expense struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}

func NewExpense(id string, name string, date time.Time, description string, amount float64) *Expense {
	return &Expense{
		ID:          id,
		Name:        name,
		Date:        date,
		Description: description,
		Amount:      amount,
	}
}
