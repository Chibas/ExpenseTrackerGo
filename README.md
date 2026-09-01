# ExpenseTrackerGo

ExpenseTrackerGo is a small command-line application for recording expenses in a local JSON file. It supports adding, listing, updating, deleting, and summarising expenses.

## Requirements

- Go 1.26.1 or later

## Installation

Clone the repository and enter the project directory:

```bash
git clone https://github.com/Chibas/ExpenseTrackerGo.git
cd ExpenseTrackerGo
```

Download the Go module dependencies:

```bash
go mod download
```

## Build

Build the executable from the repository root:

```bash
go build -o expense-tracker ./cmd
```

You can then run it as:

```bash
./expense-tracker <command> [options]
```

During development, commands can also be run without building first:

```bash
go run ./cmd <command> [options]
```

## Usage

### Add an expense

Both the description and amount are required:

```bash
./expense-tracker add --description "Lunch" --amount 12.50
```

The command prints the generated expense ID. Save this ID if you intend to update or delete the expense.

### List expenses

```bash
./expense-tracker list
```

Expenses are displayed in a table containing their ID, date, description, and amount.

### Update an expense

The ID, description, and amount are required. Replace the example ID with one returned by `add` or displayed by `list`:

```bash
./expense-tracker update \
  --id "8b889926-31f1-45c5-9f38-b1d468bd6ed6" \
  --description "Dinner" \
  --amount 24.95
```

Updating an expense also updates its recorded date to the current time.

### Delete an expense

```bash
./expense-tracker delete \
  --id "8b889926-31f1-45c5-9f38-b1d468bd6ed6"
```

### Show the total

```bash
./expense-tracker summary
```

The summary includes all stored expenses.

## Currency and amounts

ExpenseTrackerGo currently assumes all amounts are in pounds sterling (GBP). Enter amounts in pounds, optionally including pence, such as `12`, `12.50`, or `0.99`.

Internally, amounts are rounded to the nearest penny and stored as integer pence to avoid floating-point arithmetic when calculating totals. Output is formatted with the `£` symbol and two decimal places. Currency conversion and multiple currencies are not supported.

## Data location

Expenses are stored in an `expenses.json` file in the current working directory from which the command is run. For example, running the executable from the repository root creates:

```text
ExpenseTrackerGo/expenses.json
```

Run the application from the same directory each time to use the same expense data. The file is created automatically when the first expense is added and should not normally be edited by hand.

## Development checks

Run the complete test suite and static analysis from the repository root:

```bash
go test ./...
go vet ./...
```
