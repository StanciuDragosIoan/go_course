package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StanciuDragosIoan/go-course/pkg/config"
	"github.com/StanciuDragosIoan/go-course/pkg/handlers"
	"github.com/StanciuDragosIoan/go-course/pkg/render"
)

const portNumber = ":8080"

// Home is the home page handler
// func Home(w http.ResponseWriter, r *http.Request) {
	//  fmt.Fprintf(w, "This is the homepage");
// }



// addValues adds two integers and returns the sum
// func addValues(x, y int) int {
// 	return x + y
// }


//err handling
// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "cannot divide by 0")
// 		return 
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("can't divide by 0")
// 		return 0, err
// 	}
// 	result := x/y
// 	return result, nil
// }

// main is the main application function
func main() { 

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
 

	//pass reference to app
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//ignore err
	_ = http.ListenAndServe(portNumber, nil)
}