package schema

import "gopkg.in/yaml.v2"

func DecodeManifest(data []byte) (*Manifest, error) {
	m := new(Manifest)
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func EncodeManifest(m *Manifest) ([]byte, error) {
	return yaml.Marshal(&m)
}
