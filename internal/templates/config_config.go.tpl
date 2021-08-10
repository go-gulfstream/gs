package config

import (
	"bytes"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	defaultHTTPAddr = ":8080"
	defaultGRPCAddr = ":9090"
)

type internal struct {
	Addr   string `yaml:"addr"`
	Enable bool   `yaml:"enable"`
}

type kafka struct {
	Topic   string   `yaml:"topic"`
	Brokers []string `yaml:"brokers"`
}

func (k kafka) String() string {
	return strings.Join(k.Brokers, ",")
}

type http struct {
	Addr string `yml:"addr"`
}

type grpc struct {
	Addr string `yml:"addr"`
}

func parse(filename string, cfg interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(cfg); err != nil {
		return err
	}
	return nil
}