package storage

import (
	"encoding/json"
	"expanse-tracker/internal/models"
	"os"
	"path/filepath"
)

type JSONStorage struct {
	filepath string
}

func NewJSONStorage() *JSONStorage {
	return &JSONStorage{
		filepath: filepath.Join("", "expense-tracker.json"),
	}
}

func (s *JSONStorage) LoadExpenses() ([]models.Expense, error) {
	var expenses []models.Expense

	if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
		return expenses, nil
	}

	data, err := os.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func (s *JSONStorage) SaveExpenses(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filepath, data, 0644)
}
