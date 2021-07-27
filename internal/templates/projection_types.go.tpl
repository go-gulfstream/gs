package projection

import (
    "github.com/google/uuid"
)

type {{$.StreamName}} struct {
   ID        uuid.UUID        `json:"id"`
   Version   int              `json:"version"`
}