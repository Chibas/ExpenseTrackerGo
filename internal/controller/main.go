package controller

import (
	"errors"
	"flag"
	"fmt"
	"math"

	"github.com/Chibas/ExpenseTrackerGo/internal/render"
	"github.com/Chibas/ExpenseTrackerGo/internal/service"
)

type Controller interface {
	Execute(cmd Command, args []string) (err error)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Execute(cmd Command, args []string) error {
	switch cmd {
	case Add:
		addFlags := flag.NewFlagSet("add", flag.ContinueOnError)
		desc := addFlags.String("description", "", "description for your expense")
		amount := addFlags.Float64("amount", 0, "amount of your expense")
		if err := addFlags.Parse(args); err != nil {
			return err
		}
		if len(*desc) < 1 {
			return errors.New("Description can't be empty")
		}
		if *amount < 1 {
			return errors.New("Amount can't be 0")
		}
		cents := int(math.Round(*amount * 100))
		id, err := c.service.Add(*desc, cents)
		if err != nil {
			return err
		}
		fmt.Printf("\nExpense successfully added (ID: %s)", id)
		return nil
	case List:
		list, err := c.service.List()
		if err != nil {
			return err
		}
		render.Render(list)
		return nil
	case Delete:
		println("Delete")
	case Summary:
		println("Summary")
	case Update:
		println("Update")
	default:
		return errors.New("Incorrect command")
	}
	return nil
}
