package stream

import (
	"context"

	"github.com/go-gulfstream/gs/internal/source/testdata/snapshot/internal/sessioncommands"
)

func createSessionCommandController(m CommandMutation) gulfstream.ControllerFunc {
	return func(ctx context.Context, s *gulfstream.Stream, c *gulfstreamcommand.Command) (*gulfstreamcommand.Reply, error) {
		e, err := m.CreateSession(ctx, c.StreamID(), c.ID(), s.State(), c.Payload().(*sessioncommands.CreateSessionPayload))
		if err != nil {
			return c.ReplyErr(err), nil
		}
		s.Mutate(sessionevents.SessionCreated, e)
		return c.ReplyOk(s.Version()), nil
	}
}

func updateSessionCommandController(m CommandMutation) gulfstream.ControllerFunc {
	return nil
}

func removeSessionCommandController(m CommandMutation) gulfstream.ControllerFunc {
	return nil
}

func MakeCommandControllers(
	mutation CommandMutation,
	controller *gulfstream.Mutator,
) {

	controller.AddCommandController(
		sessioncommands.CreateSession,
		createSessionCommandController(mutation),
		gulfstream.WithCommandControllerCreateIfNotExists(),
	)
	controller.AddCommandController(
		sessioncommands.UpdateSession,
		updateSessionCommandController(mutation),
	)
	controller.AddCommandController(
		sessioncommands.RemoveSession,
		removeSessionCommandController(mutation),
		gulfstream.WithCommandControllerDropStream(),
	)
}
