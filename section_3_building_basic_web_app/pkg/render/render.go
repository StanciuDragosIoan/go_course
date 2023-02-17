package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/StanciuDragosIoan/go-course/pkg/config"
	"github.com/StanciuDragosIoan/go-course/pkg/models"
)

var app *config.AppConfig


func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//sets the config for the tempalte package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template;
	//disable cache for dev
	if app.UseCache {
		//get template cache from app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}



	 

	//get requested template from cache
	template, ok := templateCache[tmpl]

	if !ok {
		//kill program
		log.Fatal("Could not get templates from cache")
	}

	//create buffer (try to execute the value from the map and write it down)
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	//this is for a finer error handling
	_ = template.Execute(buf, td)

	//render template
	 _, err := buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
 }


//more advanced cache
// RenderTemplate renders templates using html
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	//create a template cache
// 	templateCache, err := CreateTemplateCache()

// 	if err != nil {
// 		//kill program
// 		log.Fatal(err)
// 	}

// 	//get requested template from cache
// 	template, ok := templateCache[tmpl]

// 	if !ok {
// 		//kill program
// 		log.Fatal(err)
// 	}

// 	//create buffer (try to execute the value from the map and write it down)
// 	buf := new(bytes.Buffer)

// 	//this is for a finer error handling
// 	_ = template.Execute(buf, nil)

// 	//render template
// 	 _, err = buf.WriteTo(w)

// 	if err != nil {
// 		log.Println(err)
// 	}
//  }

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} //empty map (same as above)

	// get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files endint with *.page.tmpl
	for _, page := range pages {
		//get file name
		name := filepath.Base(page)
		//put template with the stuff parsed from the file
		ts, err := template.New(name).ParseFiles((page))
		if err != nil {
			return myCache, err
		}

		//check for layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			//parse template (ts = template set)
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts


	}


	return myCache, nil

}



// basic cache mechanism (have to manually add the templates in createTemplateCache)

//  var templateCache = make(map[string]*template.Template);

// func RenderTemplate(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error

// 	//check if we already have the template in our cache
// 	_, inMap := templateCache[t]

// 	if!inMap {
// 		//need to create template
// 		log.Println("creating template and adding to cache")
// 		err  = createTemplateCache(t)
// 		if err !=nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = templateCache[t]

// 	err = tmpl.Execute(w,nil)
// 	if err !=nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string {
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache
// 	templateCache[t] = tmpl

// 	return nil
// }