package {{$.PackageName}}query

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/endpoint"
	transporthttp "github.com/go-kit/kit/transport/http"

	"github.com/google/uuid"
)

func NewHTTPClient(addr string) (Service, error) {
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}
	target, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	limiter := ratelimit.NewErroringLimiter(
		rate.NewLimiter(
			rate.Every(time.Second), 100),
	)

	var findEndpoint endpoint.Endpoint
	{
		findEndpoint = transporthttp.NewClient(
			http.MethodGet,
			target,
			findEncodeHTTPRequest,
			findDecodeHTTPResponse,
		).Endpoint()
		findEndpoint = limiter(findEndpoint)
		findEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(gobreaker.Settings{
				Name:        "Find",
				Timeout:     10 * time.Second,
				MaxRequests: 100,
			}))(findEndpoint)
	}

	var findOneEndpoint endpoint.Endpoint
	{
		findOneEndpoint = transporthttp.NewClient(
			http.MethodGet,
			target,
			findOneEncodeHTTPRequest,
			findOneDecodeHTTPResponse,
		).Endpoint()
		findOneEndpoint = limiter(findOneEndpoint)
		findOneEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(gobreaker.Settings{
				Name:        "FindOne",
				Timeout:     10 * time.Second,
				MaxRequests: 100,
			}))(findOneEndpoint)
	}

	return &client{
		endpoints: Endpoints{
			FindEndpoint:    findEndpoint,
			FindOneEndpoint: findOneEndpoint,
		},
	}, nil
}

type client struct {
	endpoints Endpoints
}

func (c *client) FindOne(ctx context.Context, projectionID uuid.UUID, version int) (Session, error) {
	resp, err := c.endpoints.FindOneEndpoint(ctx, FindOneRequest{
		ProjectionID: projectionID,
		Version:      version,
	})
	if err != nil {
		return Session{}, err
	}
	response := resp.(FindOneResponse)
	if len(response.Err) > 0 {
		return Session{}, errors.New(response.Err)
	}
	return response.Result, nil
}

func (c *client) Find(ctx context.Context, limit int, nextPage string, f Filter) ([]Session, string, error) {
	resp, err := c.endpoints.FindEndpoint(ctx, FindRequest{
		Limit:    limit,
		NextPage: nextPage,
		Filter:   f,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(FindResponse)
	if len(response.Err) > 0 {
		return nil, "", errors.New(response.Err)
	}
	return response.Results, response.NextPage, nil
}

func findEncodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	req := request.(FindRequest)
	r.URL.Path = "/projections"
	q := r.URL.Query()
	q.Add("limit", strconv.Itoa(req.Limit))
	q.Add("nextPage", req.NextPage)
	if req.Filter != nil {
		for k, v := range req.Filter {
			q.Add(k, v)
		}
	}
	r.URL.RawQuery = q.Encode()
	return nil
}

func findDecodeHTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp FindResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}
	if len(resp.Err) > 0 {
		return nil, errors.New(resp.Err)
	}
	return resp, nil
}

func findOneEncodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	req := request.(FindOneRequest)
	r.URL.Path = "/projections/" + req.ProjectionID.String()
	q := r.URL.Query()
	q.Add("version", strconv.Itoa(req.Version))
	r.URL.RawQuery = q.Encode()
	return nil
}

func findOneDecodeHTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp FindOneResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}
	if len(resp.Err) > 0 {
		return nil, errors.New(resp.Err)
	}
	return resp, nil
}

type FindOneRequest struct {
	ProjectionID uuid.UUID
	Version      int
}

type FindOneResponse struct {
	Err    string
	Result {{$.StreamName}}
}

type FindRequest struct {
	Limit    int
	NextPage string
	Filter   Filter
}

type FindResponse struct {
	Err      string
	NextPage string
	Results  []{{$.StreamName}}
}