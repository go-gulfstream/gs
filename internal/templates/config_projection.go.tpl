package config

type Projection struct {
	Internal internal `yaml:"internal"`
	Kafka    kafka    `yaml:"kafka"`
	GRPC     grpc     `yaml:"grpc"`
	HTTP     http     `yaml:"http"`
}

func ParseProjection(filename string) (*Projection, error) {
	cfg := &Projection{
		HTTP: http{Addr: defaultHTTPAddr},
		GRPC: grpc{Addr: defaultGRPCAddr},
	}
	if err := parse(filename, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}