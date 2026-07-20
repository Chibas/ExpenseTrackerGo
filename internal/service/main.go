package service

import (
	"time"

	"github.com/Chibas/ExpenseTrackerGo/internal/storage"
)

type Expense struct {
	ID          string
	Description string
	Date        time.Time
	Amount      float64
}

type Service interface {
	Add() error
	Delete() error
	List() error
	Update() error
}

type service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) Add() error {
	return nil
}

func (s *service) Delete() error {
	return nil
}

func (s *service) List() error {
	return nil
}

func (s *service) Update() error {
	return nil
}
