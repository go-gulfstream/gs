package projection

import (
    "strconv"
    "strings"
    "github.com/google/uuid"
)

type {{$.StreamName}} struct {
   ID        uuid.UUID
   Version   int
}

type Filter struct {
	ProjectionIDs []string
	Version       int
}

func (f Filter) String() string {
	var s string
	if len(f.ProjectionIDs) > 0 {
		s = s + "ProjectionIDs:" + strings.Join(f.ProjectionIDs, ",") + ", "
	}
	if f.Version > 0 {
		s = s + "Version:" + strconv.Itoa(f.Version) + ", "
	}
	if len(s) > 0 {
		s = "{" + s + "}"
	}
	return "Filter" + s
}