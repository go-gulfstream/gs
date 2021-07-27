package projection

import (
   "context"
   "github.com/google/uuid"
)

type Storage struct {}

func NewStorage() *Storage {
    return &Storage{}
}

func (s *Storage) Insert(ctx context.Context,  stream *{{$.StreamName}}) error {
    panic("not implemented")
    return nil
}

func (s *Storage) FindOne(ctx context.Context) (*{{$.StreamName}}, error) {
    panic("not implemented")
    return nil, nil
}

func (s *Storage) Find(ctx context.Context) ([]*{{$.StreamName}}, error) {
    panic("not implemented")
    return nil, nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
    panic("not implemented")
    return nil
}

func (s *Storage) Update(ctx context.Context, stream *{{$.StreamName}}) error {
    panic("not implemented")
    return nil
}