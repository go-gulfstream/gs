package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"{{$.GoModules}}/internal/projection"
	"github.com/google/uuid"
	"{{$.GoModules}}/pkg/{{$.PackageName}}query"
)

func MakeServerEndpoints(s {{$.PackageName}}query.Service) {{$.PackageName}}query.Endpoints {
	return {{$.PackageName}}query.Endpoints{
		FindEndpoint:    makeFindEndpoint(s),
		FindOneEndpoint: makeFindOneEndpoint(s),
	}
}

func makeFindEndpoint(s {{$.PackageName}}query.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(findRequest)
		p, err := s.Find(ctx, req.Limit, req.NextPage, req.Filter)
		return findResponse{Err: err, Results: p}, nil
	}
}

func makeFindOneEndpoint(s {{$.PackageName}}query.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(findOneRequest)
		p, err := s.FindOne(ctx, req.ProjectionID, req.Version)
		return findOneResponse{Err: err, Result: p}, nil
	}
}

type findRequest struct {
    Limit int
    NextPage string
	Filter {{$.PackageName}}query.Filter
}

type findResponse struct {
	Err     error
	Results []{{$.PackageName}}query.{{$.StreamName}}
}

func (r findResponse) error() error { return r.Err }

type findOneRequest struct {
	ProjectionID uuid.UUID
	Version      int
}

type findOneResponse struct {
	Err    error
	Result {{$.PackageName}}query.{{$.StreamName}}
}

func (r findOneResponse) error() error { return r.Err }