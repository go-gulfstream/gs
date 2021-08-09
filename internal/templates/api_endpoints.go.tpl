package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"{{$.GoModules}}/internal/projection"
	"github.com/google/uuid"
)

type Endpoints struct {
	FindOneEndpoint endpoint.Endpoint
	FindEndpoint    endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		FindEndpoint:    makeFindEndpoint(s),
		FindOneEndpoint: makeFindOneEndpoint(s),
	}
}

func makeFindEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(findRequest)
		p, err := s.Find(ctx, req.Filter)
		return findResponse{Err: err, Results: p}, nil
	}
}

func makeFindOneEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(findOneRequest)
		p, err := s.FindOne(ctx, req.ProjectionID, req.Version)
		return findOneResponse{Err: err, Result: p}, nil
	}
}

type findRequest struct {
	Filter *projection.Filter
}

type findResponse struct {
	Err     error
	Results []projection.{{$.StreamName}}
}

func (r findResponse) error() error { return r.Err }

type findOneRequest struct {
	ProjectionID uuid.UUID
	Version      int
}

type findOneResponse struct {
	Err    error
	Result projection.{{$.StreamName}}
}

func (r findOneResponse) error() error { return r.Err }