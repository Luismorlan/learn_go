package handlers

import (
	"net/http"

	"example.com/serve_html/pkg/config"
	"example.com/serve_html/pkg/models"
	"example.com/serve_html/pkg/render"
)

// Template data holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

const aboutFile = "about.page.html"
const homeFile = "home.page.html"

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, homeFile, &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again."

	render.RenderTemplate(w, aboutFile, &models.TemplateData{
		StringMap: stringMap,
	})
}
