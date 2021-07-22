package projection

type Projection interface {
    // SessionRegistered(ctx context.Context, e *event.Event) error
    // SessionUnregistered(ctx context.Context, e *event.Event) error
}

func New(
	storage *Storage,
) Projection {
	return &projection{
		storage: storage,
	}
}

type projection struct {
	storage *Storage
}