package api

import (
	"context"
	"errors"

	"{{$.GoModules}}/internal/projection"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("api: projection not found")

type Service interface {
	FindOne(ctx context.Context, projectionID uuid.UUID, version int) (projection.{{$.StreamName}}, error)
	Find(ctx context.Context, f *projection.Filter) ([]projection.{{$.StreamName}}, error)
}

type service struct {
	storage projection.Storage
}

func NewService(
	storage projection.Storage,
) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) FindOne(ctx context.Context, projectionID uuid.UUID, version int) (projection.{{$.StreamName}}, error) {
	return s.storage.FindOne(ctx, projectionID, version)
}

func (s *service) Find(ctx context.Context, f *projection.Filter) ([]projection.{{$.StreamName}}, error) {
	return s.storage.Find(ctx, f)
}