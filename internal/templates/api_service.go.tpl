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
	item, err := s.storage.FindOne(ctx, projectionID, version)
    	if err != nil {
    		return {{$.PackageName}}query.{{$.StreamName}}{}, err
    	}
    	return {{$.PackageName}}query.{{$.StreamName}}{
    		ID:      item.ID,
    		Version: item.Version,
    	}, nil
}

func (s *Service) Find(ctx context.Context, limit int, nextPage string, f {{$.PackageName}}query.Filter) ([]{{$.PackageName}}query.{{$.StreamName}}, string, error) {
	filter := projection.Filter{
	   Limit: limit,
	   NextPage: nextPage,
	}
	items, nextPage, err := s.storage.Find(ctx, filter)
	if err != nil {
	     return nil, "", err
	}
    results := make([]{{$.PackageName}}query.{{$.StreamName}}, len(items))
    for i, r := range items {
       results[i] = {{$.PackageName}}query.{{$.StreamName}}{
          ID: r.ID,
          Version: r.Version,
       }
    }
	return results, nextPage, nil
}