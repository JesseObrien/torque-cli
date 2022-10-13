package scaffold

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestScaffoldEntireFileStructure(t *testing.T) {
	testFS := afero.NewMemMapFs()

	appName := "TestApp"

	conf := ScaffoldConfig{
		FS:      testFS,
		AppName: appName,
		ModName: "testmod",
	}

	s := NewScaffolder(conf)

	err := s.Scaffold()

	for _, file := range ProjectFiles {
		_, err := testFS.Stat(fmt.Sprintf("/%s", file.outputFilePath))
		assert.NoError(t, err)
	}

	assert.NoError(t, err)
}
