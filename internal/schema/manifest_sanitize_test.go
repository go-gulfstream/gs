package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeManifest_PackageName(t *testing.T) {
	manifest := new(Manifest)
	manifest.PackageName = "My Package Name"
	SanitizeManifest(manifest)
	assert.Equal(t, "mypackagename", manifest.PackageName)
}

func TestSanitizeManifest_StreamName(t *testing.T) {
	manifest := new(Manifest)
	manifest.StreamName = "my    Stream Name 13123"
	SanitizeManifest(manifest)
	assert.Equal(t, "Mystreamname13123", manifest.StreamName)
}
