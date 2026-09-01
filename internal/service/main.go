package service

import (
	"errors"
	"slices"
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
	Delete(id string) (bool, error)
	List() (data []Expense, err error)
	Update(id string, description string, amount int) (bool, error)
	Summary() (int, error)
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

func (s *service) Delete(id string) (bool, error) {
	expenses, err := s.storage.Read()
	if err != nil {
		return false, err
	}
	i := slices.IndexFunc(expenses, func(e Expense) bool {
		return e.ID == id
	})
	if i == -1 {
		return false, nil
	}
	modifiedExpenses := slices.Delete(expenses, i, i+1)
	if err := s.storage.Write(modifiedExpenses); err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) List() (data []Expense, err error) {
	return s.storage.Read()
}

func (s *service) Update(id string, description string, amount int) (bool, error) {
	expenses, err := s.storage.Read()
	if err != nil {
		return false, err
	}
	if len(id) < 1 {
		return false, errors.New("id can't be empty")
	}
	if len(description) < 1 {
		return false, errors.New("Description can't be empty")
	}

	if amount < 1 {
		return false, errors.New("Amount can't be less than 1")
	}
	i := slices.IndexFunc(expenses, func(e Expense) bool {
		return e.ID == id
	})
	if i == -1 {
		return false, nil
	}
	expense := expenses[i]
	expense.Description = description
	expense.Date = time.Now()
	expense.Amount = amount

	expenses[i] = expense

	if err := s.storage.Write(expenses); err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) Summary() (sum int, err error) {
	expenses, err := s.storage.Read()
	if err != nil {
		return 0, err
	}
	for _, expense := range expenses {
		sum += expense.Amount
	}
	return sum, nil
}
