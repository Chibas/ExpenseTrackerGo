package main

import (
	"flag"
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
	}
	cmd := args[0]

	if len(args) > 0 {
		fmt.Println("command:", args[0])
		args = args[1:] // remove the command word
	}

	fs := flag.NewFlagSet("expense-tracker", flag.ContinueOnError)
	description := fs.String("description", "", "Description for your expense")
	amount := fs.Float64("amount", 0, "amount")

	if err := fs.Parse(args); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Flag description", *description, "amount", *amount)

	storage := storage.NewStorage("expenses")
	svc := service.NewService(storage)
	ctrl := controller.NewController(svc)
	ctrl.Execute(controller.Command(cmd))
}
