package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

    "github.com/google/uuid"
	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"

	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

var errBadRouting = errors.New("api: inconsistent mapping between route and handler")

func MakeHTTPHandler(s {{$.PackageName}}query.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []transporthttp.ServerOption{
		transporthttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		transporthttp.ServerErrorEncoder(encodeError),
	}
	r.Methods(http.MethodGet).Path("/projections/{projection_id}").Handler(transporthttp.NewServer(
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
	vars := mux.Vars(r)
	pstr, ok := vars["projection_id"]
	if !ok {
		return nil, errBadRouting
	}
	id, err := uuid.Parse(pstr)
	if err != nil {
		return nil, errBadRouting
	}
	req := findOneRequest{ProjectionID: id}
	vstr, ok := vars["version"]
	if ok {
		req.Version, err = strconv.Atoi(vstr)
		if err != nil {
			return nil, errBadRouting
		}
	}
	return req, nil
}

func decodeFindRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return findRequest{
		Filter: map[string]string{},
	}, nil
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
	case errNotFound:
		return http.StatusNotFound
	case errBadRouting:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}