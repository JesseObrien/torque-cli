package scaffold

import (
	"embed"

	"github.com/spf13/afero"
)

//go:embed templates/*
var templateFS embed.FS

type ScaffoldConfig struct {
	// The filesystem to use
	FS afero.Fs

	// Application Name
	AppName string
	ModName string

	// Default controller name
	ControllerName string

	// Custom specified directory for the project files
	Path string

	ORM   bool
	AWS   bool
	Redis bool
}

type Scaffolder struct {
	Config ScaffoldConfig
}

func NewScaffolder(config ScaffoldConfig) *Scaffolder {
	if config.FS == nil {
		config.FS = afero.NewOsFs()
	}

	if config.ControllerName == "" {
		config.ControllerName = "app"
	}

	return &Scaffolder{Config: config}
}

func (s *Scaffolder) ScaffoldNewProject() error {
	for _, file := range ProjectFiles {
		if err := file.ScaffoldTemplate(s.Config.FS, s.Config); err != nil {
			return err
		}
	}

	return nil
}
