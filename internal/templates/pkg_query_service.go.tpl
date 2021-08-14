package {{$.PackageName}}query

import (
	"context"
	"github.com/google/uuid"
	"github.com/go-kit/kit/endpoint"
	"strconv"
	"strings"
)

type Service interface {
	FindOne(ctx context.Context, projectionID uuid.UUID, version int) ({{$.StreamName}}, error)
	Find(ctx context.Context, limit int, nextPage string, f Filter) ([]{{$.StreamName}}, error)
}

type Middleware func(Service) Service

type {{$.StreamName}} struct {
	ID      uuid.UUID
	Version int
}

type Endpoints struct {
	FindOneEndpoint endpoint.Endpoint
	FindEndpoint    endpoint.Endpoint
}

type Filter map[string]string