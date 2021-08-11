package schema

const (
	HTTPCommandBus commandBusAdapter = iota
	GRPCCommandBus
	NATSCommandBus
)

const (
	HTTPCommandBusName = "http"
	GRPCCommandBusName = "grpc"
	NATSCommandBusName = "nats"
)

var CommandBusAdapters = map[commandBusAdapter]string{
	NATSCommandBus: NATSCommandBus.String(),
	GRPCCommandBus: GRPCCommandBus.String(),
	HTTPCommandBus: HTTPCommandBus.String(),
}

type commandBusAdapter int

func (a commandBusAdapter) String() string {
	switch a {
	case HTTPCommandBus:
		return "HTTP"
	case GRPCCommandBus:
		return "GRPC"
	case NATSCommandBus:
		return "NATS"
	}
	return "Unknown"
}
