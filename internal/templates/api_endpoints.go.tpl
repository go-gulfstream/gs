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
		req := request.({{$.PackageName}}query.FindRequest)
		p, np, err := s.Find(ctx, req.Limit, req.NextPage, req.Filter)
		return {{$.PackageName}}query.FindResponse{Err: err2str(err),  NextPage: np, Results: p}, nil
	}
}

func makeFindOneEndpoint(s {{$.PackageName}}query.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.({{$.PackageName}}query.FindOneRequest)
		p, err := s.FindOne(ctx, req.ProjectionID, req.Version)
		return {{$.PackageName}}query.FindOneResponse{Err: err2str(err), Result: p}, nil
	}
}

func err2str(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}