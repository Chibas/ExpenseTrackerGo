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
		deleteFlags := flag.NewFlagSet("delete", flag.ContinueOnError)
		id := deleteFlags.String("id", "", "ID of your expense to remove")
		if err := deleteFlags.Parse(args); err != nil {
			return err
		}
		if len(*id) < 1 {
			return errors.New("ID can't be empty")
		}
		deleted, err := c.service.Delete(*id)
		if err != nil {
			return err
		}
		if deleted {
			fmt.Printf("\nExpense (ID: %s) successfully deleted ", *id)
		} else {
			return fmt.Errorf("expense with ID %q not found", *id)
		}
		return nil

	case Summary:
		sum, err := c.service.Summary()
		if err != nil {
			return fmt.Errorf("Summary failed %w", err)
		}
		fmt.Printf("Total expenses %s\n", render.FormatAmount(sum))
		return nil
	case Update:
		updateFlags := flag.NewFlagSet("update", flag.ContinueOnError)
		id := updateFlags.String("id", "", "ID of your expense")
		desc := updateFlags.String("description", "", "description for your expense")
		amount := updateFlags.Float64("amount", 0, "amount of your expense")
		if err := updateFlags.Parse(args); err != nil {
			return err
		}
		if len(*desc) < 1 {
			return errors.New("Description can't be empty")
		}
		if *amount < 1 {
			return errors.New("Amount can't be 0")
		}
		cents := int(math.Round(*amount * 100))
		updated, err := c.service.Update(*id, *desc, cents)
		if err != nil {
			return err
		}
		if updated {
			fmt.Printf("\nExpense (ID: %s) successfully updated ", *id)
		} else {
			return fmt.Errorf("expense with ID %q not found", *id)
		}
		return nil

	default:
		return errors.New("Incorrect command")
	}
}
