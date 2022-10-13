package scaffold

import (
	"bufio"
	"fmt"
	"text/template"

	"github.com/spf13/afero"
)

const (
	GITKEEP   = ".gitkeep"
	GITIGNORE = ".gitignore"

	ENVFILE        = ".env"
	DOCKERFILE     = "dockerfile"
	DOCKER_COMPOSE = "docker-compose"
	TORQUE_CONFIG  = "torque-config"
	MODD_CONFIG    = "modd-config"

	MAIN_FILE    = "main-file"
	HTTP_SERVICE = "http-service"
	CONTROLLER   = "controller"
)

type ProjectFile struct {
	templateFilePath string
	outputFilePath   string
}

func (p *ProjectFile) ScaffoldCustomTemplate(fs afero.Fs, customPath string, data any) error {
	p.outputFilePath = customPath

	return p.ScaffoldTemplate(fs, data)
}

func (p *ProjectFile) ScaffoldTemplate(fs afero.Fs, data any) error {
	info, err := fs.Stat(p.outputFilePath)

	if err == nil || info != nil {
		return fmt.Errorf("File already exists at path %s, I will not overwrite it!", p.outputFilePath)
	}

	f, err := fs.Create(p.outputFilePath)

	if err != nil {
		return err
	}

	defer f.Close()

	fmt.Printf("Scaffolding %s\n", p.outputFilePath)
	tmp := template.Must(template.ParseFS(templateFS, fmt.Sprintf("templates/%s", p.templateFilePath)))

	buf := bufio.NewWriter(f)
	defer buf.Flush()
	if err := tmp.Execute(buf, data); err != nil {
		return err
	}

	return nil
}

var ProjectFiles = map[string]ProjectFile{
	// Git helpers
	GITIGNORE: {".gitignore", ".gitignore"},
	GITKEEP:   {".gitkeep", "dist/.gitkeep"},

	// Config Files
	ENVFILE:        {"config/env.example.go.tmpl", ".env.example"},
	DOCKERFILE:     {"config/Dockerfile.tmpl", "Dockerfile"},
	DOCKER_COMPOSE: {"config/docker-compose.yml.tmpl", "docker-compose.yml"},
	TORQUE_CONFIG:  {"config/torque.yml.tmpl", "torque.yml"},
	MODD_CONFIG:    {"config/modd.conf.tmpl", "modd.conf"},

	// Main project files
	MAIN_FILE:    {"main/main.go.tmpl", "cmd/main/main.go"},
	HTTP_SERVICE: {"http/http.go.tmpl", "internal/http/http.go"},
	CONTROLLER:   {"http/controller.go.tmpl", "internal/http/controllers/%s.go"},
}
