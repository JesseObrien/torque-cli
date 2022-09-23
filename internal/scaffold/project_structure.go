package scaffold

type ProjectFile struct {
	templateFilePath string
	outputFilePath   string
}

var ProjectFiles = []ProjectFile{
	// Git helpers
	{".gitignore", ".gitignore"},
	{".gitkeep", "dist/.gitkeep"},

	// Config Files
	{"config/env.example.go.tmpl", ".env.example"},
	{"config/Dockerfile.tmpl", "Dockerfile"},
	{"config/docker-compose.yml.tmpl", "docker-compose.yml"},
	{"config/torque.yml.tmpl", "torque.yml"},

	// Main project files
	{"main/main.go.tmpl", "cmd/main/main.go"},
	{"http/http.go.tmpl", "internal/http/http.go"},
}
