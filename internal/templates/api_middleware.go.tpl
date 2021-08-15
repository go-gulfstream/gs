package api

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"{{$.GoModules}}/internal/projection"
	"{{$.GoModules}}/pkg/{{$.PackageName}}query"
	"github.com/google/uuid"
)

type loggingMiddleware struct {
	logger log.Logger
	next   {{$.PackageName}}query.Service
}

func LoggingMiddleware(logger log.Logger) {{$.PackageName}}query.Middleware {
	return func(next {{$.PackageName}}query.Service) {{$.PackageName}}query.Service {
		return loggingMiddleware{logger, next}
	}
}

func (m loggingMiddleware) FindOne(ctx context.Context, projectionID uuid.UUID, version int) (p {{$.PackageName}}query.{{$.StreamName}}, err error) {
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

func (m loggingMiddleware) Find(ctx context.Context, limit int, nextPage string, f {{$.PackageName}}query.Filter) (p []{{$.PackageName}}query.{{$.StreamName}}, np string, err error) {
	defer func(startTime time.Time) {
		m.logger.Log(
			"method", "Find",
			"filter", f,
			"found", len(p),
			"nextPage", np,
			"took", time.Since(startTime),
			"err", err)
	}(time.Now())
	p, np, err = m.next.Find(ctx, limit, nextPage, f)
	return
}