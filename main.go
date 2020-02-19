package main

import (
	"os"

	"github.com/unrolled/render"

	log "github.com/sirupsen/logrus"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)
	
type response struct {
	ID uint `json:"id,omitempty"`
}

func main() {
	// set log level
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// load config file
	path := "config.yml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	cfg := &DefaultConfig
	cfg.Load(path)

	// connect to DB
	dbc := NewDB(cfg)
	defer dbc.Close()

	// render subsystem
	render := render.New()

	// configure server
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	

	// CORS
	cware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:         300,
	})
	r.Use(cware.Handler)

	apiRoutes(r, dbc, render, cfg)

	// start server
	log.Printf("Hearing at %s%s", cfg.Server.Port, cfg.Server.Path)
	http.ListenAndServe(":"+cfg.Server.Port, r)
}
