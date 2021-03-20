package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/handlers"
	"github.com/tsawler/go-course/pkg/render"
)

var port = "8080"

// if the first letter of a function is lowercase
// the function is not accessible elsewhere from outside the package

// if its uppercase, it can be accessible from outside the package

/*
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes writted: %d", n))
	})
*/

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change to true if in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port " + port)
	// _ = http.ListenAndServe(":"+port, nil)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
