package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"

	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []transporthttp.ServerOption{
		transporthttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		transporthttp.ServerErrorEncoder(encodeError),
	}
	r.Methods(http.MethodGet).Path("/projection/{projection_id}").Handler(transporthttp.NewServer(
		e.FindOneEndpoint,
		decodeFindOneRequest,
		encodeResponse,
		options...,
	))
	r.Methods(http.MethodGet).Path("/projections").Handler(transporthttp.NewServer(
		e.FindEndpoint,
		decodeFindRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeFindOneRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
    // TODO:
    panic("TODO")
	return nil, nil
}

func decodeFindRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	// TODO:
	panic("TODO")
	return nil, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		err = errors.New("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}