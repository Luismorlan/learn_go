package main

import (
	"log"
	"net/http"

	"example.com/serve_html/pkg/config"
	"example.com/serve_html/pkg/handlers"
	"example.com/serve_html/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplaceCache = tc
	app.UseCache = false

	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Starting server on port:", port)
	http.ListenAndServe(port, nil)
}
