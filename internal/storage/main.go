package storage

type Storage interface {
	ReadFile() error
	WriteFile(data any) error
}

type storage struct {
	fileName string
}

func NewStorage(name string) Storage {
	return &storage{
		fileName: name,
	}
}

func (s *storage) ReadFile() error {
	return nil
}

func (s *storage) WriteFile(data any) error {
	return nil
}
