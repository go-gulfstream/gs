package projection

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Insert(ctx context.Context, stream {{$.StreamName}}) error
	FindOne(ctx context.Context, id uuid.UUID, version int) ({{$.StreamName}}, error)
	Find(ctx context.Context, f *Filter) ([]{{$.StreamName}}, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, stream {{$.StreamName}}) error
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) Insert(ctx context.Context, stream {{$.StreamName}}) error {
	return nil
}

func (s *storage) FindOne(ctx context.Context, id uuid.UUID, version int) ({{$.StreamName}}, error) {
	return {{$.StreamName}}{
        ID:      id,
        Version: version,
	}, nil
}

func (s *storage) Find(ctx context.Context, f *Filter) ([]{{$.StreamName}}, error) {
	return []{{$.StreamName}}{}, nil
}

func (s *storage) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (s *storage) Update(ctx context.Context, stream {{$.StreamName}}) error {
	return nil
}