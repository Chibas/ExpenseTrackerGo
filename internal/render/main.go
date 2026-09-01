package render

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Chibas/ExpenseTrackerGo/internal/service"
	"github.com/aquasecurity/table"
)

func Render(expenses []service.Expense) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "ID", "Date", "Description", "Amount")

	for i, expense := range expenses {
		table.AddRow(
			strconv.Itoa(i+1),
			expense.ID,
			expense.Date.Format("2006-01-02 15:04:05"),
			expense.Description,
			FormatAmount(expense.Amount),
		)
	}

	table.Render()
}

func FormatAmount(amount int) string {
	return fmt.Sprintf("£%d.%02d", amount/100, amount%100)
}
