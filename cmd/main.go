package main

import (
	"fmt"
	"os"

	"github.com/Chibas/ExpenseTrackerGo/internal/controller"
	"github.com/Chibas/ExpenseTrackerGo/internal/service"
	"github.com/Chibas/ExpenseTrackerGo/internal/storage"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("No Command provided")
		return
	}
	cmd := args[0]

	if len(args) > 0 {
		fmt.Println("command:", args[0])
		args = args[1:] // remove the command word
	}

	storage := storage.NewStorage[[]service.Expense]("expenses.json")
	svc := service.NewService(storage)
	ctrl := controller.NewController(svc)
	err := ctrl.Execute(controller.Command(cmd), args)
	if err != nil {
		fmt.Printf("An error occurred %s", err)
		os.Exit(2)
	}
}
