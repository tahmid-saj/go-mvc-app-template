package handlers

import (
	// "fmt"
	"net/http"
	// "html/template"
	"github.com/tahmidsaj/go-course/pkg/config"
	"github.com/tahmidsaj/go-course/pkg/models"
	"github.com/tahmidsaj/go-course/pkg/render"
)

// Repository var used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the new repository
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
	// fmt.Fprint(w, "This is the home page")

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprint(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send the data to the template

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}


