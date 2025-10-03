package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nktshklv-dev/storage-go/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.files[newFile.Id] = newFile

	return newFile, nil
}

func (s Storage) GetById(id uuid.UUID) (*file.File, error) {
	gotFile, ok := s.files[id]

	if !ok {
		return nil, fmt.Errorf("file with id %d was not found", id)
	}

	return gotFile, nil
}
