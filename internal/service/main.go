package service

import (
	"errors"
	"time"

	"github.com/Chibas/ExpenseTrackerGo/internal/storage"
	"github.com/google/uuid"
)

type Expense struct {
	ID          string
	Description string
	Date        time.Time
	Amount      int
}

type Service interface {
	Add(description string, amount int) (id string, err error)
	Delete() error
	List() (data []Expense, err error)
	Update() error
}

type service struct {
	storage storage.Storage[[]Expense]
}

func NewService(storage storage.Storage[[]Expense]) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) Add(description string, amount int) (id string, err error) {
	var expense Expense
	expenses, err := s.storage.Read()
	if err != nil {
		return "", err
	}

	if len(description) < 1 {
		return "", errors.New("Description can't be empty")
	}

	if amount < 1 {
		return "", errors.New("Amount can't be less than 1")
	}

	expense = Expense{
		ID:          uuid.New().String(),
		Description: description,
		Date:        time.Now(),
		Amount:      amount,
	}

	expenses = append(expenses, expense)
	return expense.ID, s.storage.Write(expenses)
}

func (s *service) Delete() error {
	return nil
}

func (s *service) List() (data []Expense, err error) {
	return s.storage.Read()
}

func (s *service) Update() error {
	return nil
}
