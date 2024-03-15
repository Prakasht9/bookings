package handlers

import (
	"net/http"

	"github.com/prakasht9/bookings/pkg/config"
	"github.com/prakasht9/bookings/pkg/models"
	"github.com/prakasht9/bookings/pkg/render"
)

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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, World")
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perfrom some business logic and send data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remove_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
