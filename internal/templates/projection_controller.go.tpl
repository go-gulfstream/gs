package projection

import (
	"context"

	"github.com/go-gulfstream/gulfstream/pkg/event"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"{{$.Project.GoModules}}/{{$.Project.Name}}/pkg/events"
)

func NewController(p Projection) *gulfstream.Projection {
	projection := gulfstream.NewProjection()

    // For example:
	// projection.AddEventController(
	//	events.SessionRegistered,
	//	sessionRegisteredController(p),
	// )

    // For example:
	// projection.AddEventController(
	//	events.SessionUnregisteredEvent,
	//	sessionUnregisteredController(p),
	// )

	return projection
}

// func sessionRegisteredController(p Projection) gulfstream.EventHandlerFunc {
//	return func(ctx context.Context, e *event.Event) error {
//		return p.SessionRegistered(ctx, e)
//	}
// }
//
// func sessionUnregisteredController(p Projection) gulfstream.EventHandlerFunc {
//	return func(ctx context.Context, e *event.Event) error {
//		return p.SessionUnregistered(ctx, e)
//	}
// }