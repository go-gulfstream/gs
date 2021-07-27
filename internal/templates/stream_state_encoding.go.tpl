package stream

import (
   "encoding/json"
)

type Snapshot struct {}

func (s *root) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *root) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}