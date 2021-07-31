package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifact_ParseCommandMutationInterface(t *testing.T) {
	snapshot, err := NewSnapshot("./testdata/snapshot")
	assert.Nil(t, err)
	artifact := NewArtifact(snapshot)
	assert.Equal(t, parseBreak, artifact.parseCommandMutationInterface())
	assert.Len(t, artifact.commandMutations, 3)
}

func TestArtifact_ParseCommandController(t *testing.T) {
	snapshot, err := NewSnapshot("./testdata/snapshot")
	assert.Nil(t, err)
	artifact := NewArtifact(snapshot)
	assert.Equal(t, parseBreak, artifact.parseCommandMutationInterface())
	assert.Nil(t, artifact.parseCommandController())
	assert.Len(t, artifact.commandController, 3)
}

func TestArtifact_ParseMakeCommandControllers(t *testing.T) {
	snapshot, err := NewSnapshot("./testdata/snapshot")
	assert.Nil(t, err)
	artifact := NewArtifact(snapshot)
	assert.Equal(t, parseBreak, artifact.parseMakeCommandControllers())
	assert.NotNil(t, artifact.makeCommandControllers)
}
