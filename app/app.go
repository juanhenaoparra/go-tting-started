package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/juanhenaoparra/go-tting-started/models"
)

// App is the main application for http server
type App struct {
	Router      *chi.Mux
	ActiveQueue *models.MessageQueue
}

func (a *App) setupRouter() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	a.Router = r
}

// Start starts the http server
func (a App) Start(port int) error {
	return http.ListenAndServe("127.0.0.1:"+fmt.Sprint(port), a.Router)
}

// NewApp creates a new App
func NewApp() (*App, error) {
	app := &App{}
	app.setupRouter()
	return app, nil
}
