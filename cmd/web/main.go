package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/tahmidsaj/go-course/pkg/render"
	"github.com/tahmidsaj/go-course/pkg/config"
	"github.com/tahmidsaj/go-course/pkg/handlers"
	// "errors"
	// "html/template"
)

const portNumber = ":8080"



// addValues adds two integers and return the sum
// func addValues(x, y int) int {
// 	return x + y
// }

func Divide(w http.ResponseWriter, r *http.Request) {
	// f, err := divideValues(100.0, 0.0)

	// if err != nil {
	// 	fmt.Fprintf(w, "Cannot divide by 0")
	// 	return
	// }

	// fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("Cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

// main is the main application function
func main() {
	// fmt.Println("Hello world");

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	// })

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}