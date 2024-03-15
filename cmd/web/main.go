package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prakasht9/bookings/pkg/config"
	"github.com/prakasht9/bookings/pkg/handlers"
	"github.com/prakasht9/bookings/pkg/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates((&app))

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// // http.HandleFunc("/divide", Divide)

	// fmt.Printf("Starting application on port: %s \n", port_number)

	// _ = http.ListenAndServe(port_number, nil)
}
