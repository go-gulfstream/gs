package api

import (
	"context"
	"errors"

    "{{$.GoModules}}/pkg/{{$.PackageName}}query"
	"{{$.GoModules}}/internal/projection"
	"github.com/google/uuid"
)

var errNotFound = errors.New("api: projection not found")

var _ {{$.PackageName}}query.Service = (*Service)(nil)

type Service struct {
	storage projection.Storage
}

func NewService(
	storage projection.Storage,
) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) FindOne(ctx context.Context, projectionID uuid.UUID, version int) ({{$.PackageName}}query.{{$.StreamName}}, error) {
	return {{$.PackageName}}query.{{$.StreamName}}{}, nil
}

func (s *Service) Find(ctx context.Context, limit int, nextPage string, f {{$.PackageName}}query.Filter) ([]{{$.PackageName}}query.{{$.StreamName}}, string, error) {
	return []{{$.PackageName}}query.{{$.StreamName}}{}, "", nil
}