package scaffold

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type ScaffoldConfig struct {
	AppName string
}

type Scaffolder struct {
	Config ScaffoldConfig
}

func NewScaffolder(config ScaffoldConfig) *Scaffolder {
	return &Scaffolder{Config: config}
}

func (s *Scaffolder) Scaffold() error {
	if err := s.scaffoldTemplate("main/main.go.tmpl", "cmd/main/main.go", nil); err != nil {
		return err
	}
	return nil
}

func (s *Scaffolder) scaffoldTemplate(tmpl string, file string, data any) error {
	f, err := os.Create(file)
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
