package config

type Stream struct {
}

func ParseStream(filename string) (*Stream, error) {
	cfg := &Stream{}
	if err := parse(filename, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}