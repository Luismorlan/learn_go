package config

import (
	"html/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplaceCache map[string]*template.Template
}
