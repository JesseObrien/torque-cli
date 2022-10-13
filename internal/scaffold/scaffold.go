package scaffold

import (
	"bufio"
	"embed"
	"fmt"
	"text/template"

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
	return &Scaffolder{Config: config}
}

func (s *Scaffolder) Scaffold() error {
	for _, file := range ProjectFiles {
		if err := s.scaffoldTemplate(file.templateFilePath, file.outputFilePath, s.Config); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scaffolder) scaffoldTemplate(tmpl string, file string, data any) error {
	fullPath := fmt.Sprintf("%s/%s", s.Config.Path, file)
	f, err := s.Config.FS.Create(fullPath)
	defer f.Close()

	if err != nil {
		return err
	}

	tmp := template.Must(template.ParseFS(templateFS, fmt.Sprintf("templates/%s", tmpl)))

	buf := bufio.NewWriter(f)
	defer buf.Flush()
	if err := tmp.Execute(buf, data); err != nil {
		return err
	}

	return nil
}
