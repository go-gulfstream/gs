package api

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/go-kit/log"
	"{{$.GoModules}}/internal/projection"
	"github.com/google/uuid"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

func (m loggingMiddleware) FindOne(ctx context.Context, projectionID uuid.UUID, version int) (p projection.{{$.StreamName}}, err error) {
	defer func(startTime time.Time) {
		m.logger.Log(
			"method", "FindOne",
			"projectionID", projectionID,
			"version", version,
			"took", time.Since(startTime),
			"err", err)
	}(time.Now())
	p, err = m.next.FindOne(ctx, projectionID, version)
	return
}

func (m loggingMiddleware) Find(ctx context.Context, f *projection.Filter) (p []projection.{{$.StreamName}}, err error) {
	defer func(startTime time.Time) {
		m.logger.Log(
			"method", "Find",
			"filter", f,
			"found", len(p),
			"took", time.Since(startTime),
			"err", err)
	}(time.Now())
	p, err = m.next.Find(ctx, f)
	return
}

type metricMiddleware struct {
	logger               log.Logger
	next                 Service
	findOneCounter       metrics.Counter
	findOneHistogram     metrics.Histogram
	findCounter          metrics.Counter
	findCounterHistogram metrics.Histogram
}