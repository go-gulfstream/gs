package config

type Stream struct {
 InternalHTTP http `yaml:"internal_http"`
 {{if $.StreamStorage.IsPostgres -}}
 Postgres postgres `yaml:"postgres"`
 {{else if $.StreamStorage.IsRedis -}}
 Redis redis `yaml:"redis"`
 {{end -}}
 {{if $.StreamPublisher.IsKafka -}}
 Kafka kafka `yaml:"kafka"`
 {{end}}
  {{if $.CommandBus.IsGRPC -}}
  CommandBusGRPC grpc `yaml:"commandbus_grpc"`
  {{else if $.CommandBus.IsNATS -}}
  CommandBusNATS nats `yaml:"commandbus_nats"`
  {{else if $.CommandBus.IsHTTP -}}
  CommandBusHTTP http `yaml:"commandbus_http"`
  {{end}}
}

func ParseStream(filename string) (*Stream, error) {
	cfg := &Stream{}
	if err := parse(filename, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}