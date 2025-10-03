package file

import "github.com/google/uuid"

type File struct {
	Name string
	Id   uuid.UUID
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileId, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return &File{
		Name: filename,
		Id:   fileId,
		Data: blob,
	}, nil
}
