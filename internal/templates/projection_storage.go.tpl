package projection

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Insert(ctx context.Context, stream *Session) error
	FindOne(ctx context.Context, id uuid.UUID, version int) (*Session, error)
	Find(ctx context.Context, f *Filter) ([]*Session, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, stream *Session) error
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) Insert(ctx context.Context, stream *Session) error {
	panic("not implemented")
	return nil
}

func (s *storage) FindOne(ctx context.Context, id uuid.UUID, version int) (*Session, error) {
	panic("not implemented")
	return nil, nil
}

func (s *storage) Find(ctx context.Context, f *Filter) ([]*Session, error) {
	panic("not implemented")
	return nil, nil
}

func (s *storage) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
	return nil
}

func (s *storage) Update(ctx context.Context, stream *Session) error {
	panic("not implemented")
	return nil
}