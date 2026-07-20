package controller

import "github.com/Chibas/ExpenseTrackerGo/internal/service"

type Controller interface {
	Execute(Command) error
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Execute(cmd Command) error {
	switch cmd {
	case Add:
		println("Add")
	case List:
		println("List")
	case Delete:
		println("Delete")
	case Summary:
		println("Summary")
	default:
		println("\nIncorrect command")
	}
	return nil
}
