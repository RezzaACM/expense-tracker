# Expense Tracker CLI

A command-line tool for managing personal expenses written in Go. Sample solution for the [expense-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Features

- Add new expenses with name, amount, and description
- List all expenses in a tabular format
- Delete expenses by ID
- Update existing expenses
- View expense summaries by month or total

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/arikchakma/backend-projects.git
cd expense-tracker
```

## Installation

```bash
go install expanse-tracker

Usage
# Add an expense
expanse-tracker add --name "Groceries" --amount 50.50 --description "Weekly groceries"

# List all expenses
expanse-tracker list

# Delete an expense
expanse-tracker delete --id <expense-id>

# Update an expense
expanse-tracker update --id <expense-id> --name "New Name" --amount 75.00

# View expense summary
expanse-tracker summary           #* View total expenses*
expanse-tracker summary --month 1 #* View January expenses*
```

### Project Structure

    .cmd
        ├── add.go              # Add new expenses
        ├── delete.go           # Delete expenses
        ├── list.go             # List expenses
        ├── update.go           # Update expenses
        └── summary.go          # View summaries
    .internal
        ├── models              # Data models
            └── expense.go          # Expense model
        ├── storage             # JSON storage implementation
        └── utils               # Utility functions
            └── util.go             # Util

Data Storage
Expenses are stored in a local JSON file (expense-tracker.json) in the current directory.

### Dependencies

    github.com/spf13/cobra - CLI framework
    github.com/google/uuid - UUID generation

### License

MIT License
