package handlers

import (
	"net/http"

	"github.com/StanciuDragosIoan/go-course/pkg/config"
	"github.com/StanciuDragosIoan/go-course/pkg/models"
	"github.com/StanciuDragosIoan/go-course/pkg/render"
)

//Repo (the repository used by the handlers)
var Repo *Repository

//Repository (is the repository type)
type Repository struct {
	App *config.AppConfig
}

// NewRepo (creates new repository)
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
			App: a,
	}
}


// NewHandlers (sets the repositroy for the handlers)
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	//perform logic get data

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello test map!"

	//inject data in template
   render.RenderTemplate(w, "about.page.tmpl",  &models.TemplateData{
	StringMap: stringMap,
   })
}

