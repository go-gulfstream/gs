package schema

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func DecodeManifest(data []byte) (*Manifest, error) {
	m := new(Manifest)
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("invalid manifest file")
	}
	return m, nil
}

func EncodeManifest(m *Manifest) ([]byte, error) {
	return yaml.Marshal(&m)
}
