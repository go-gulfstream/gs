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