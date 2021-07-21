package projection

import (
    "github.com/google/uuid"
)

type Stream struct {
   ID        uuid.UUID        `json:"id"`
   Version   int              `json:"version"`
}