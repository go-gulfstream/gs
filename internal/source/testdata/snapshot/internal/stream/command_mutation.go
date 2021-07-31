package stream

import (
	"context"

	"github.com/go-gulfstream/gs/internal/source/testdata/snapshot/internal/sessionevents"

	"github.com/go-gulfstream/gs/internal/source/testdata/snapshot/internal/sessioncommands"

	"github.com/google/uuid"
)

type CommandMutation interface {
	CreateSession(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *sessioncommands.CreateSessionPayload) (*sessionevents.SessionCreatedPayload, error)
	UpdateSession(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *sessioncommands.UpdateSessionPayload) (*sessionevents.SessionUpdatedPayload, error)
	RemoveSession(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *sessioncommands.RemoveSessionPayload) (*sessionevents.SessionRemovedPayload, error)
}

type State interface {
}
