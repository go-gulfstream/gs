package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSnapshot(t *testing.T) {
	snapshot, err := NewSnapshot("./testdata/snapshot")
	assert.Nil(t, err)
	assert.NotNil(t, snapshot)
	assert.NotEmpty(t, snapshot.TotalFiles())
	err = snapshot.Walk(func(info FileInfo) error {
		assert.Contains(t, info.Path(), ".go")
		return nil
	})
	assert.Nil(t, err)
}
