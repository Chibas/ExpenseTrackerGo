package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type Storage[T any] interface {
	Read() (data T, err error)
	Write(data T) error
}

type storage[T any] struct {
	fileName string
}

func NewStorage[T any](name string) Storage[T] {
	return &storage[T]{
		fileName: name,
	}
}

func (s *storage[T]) Read() (data T, err error) {
	var result T
	fileData, err := os.ReadFile(s.fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return result, nil
		}
		return result, err
	}

	err = json.Unmarshal(fileData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *storage[T]) Write(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.fileName, fileData, 0644)
}
