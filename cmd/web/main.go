package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prakasht9/bookings/internal/config"
	"github.com/prakasht9/bookings/internal/handlers"
	"github.com/prakasht9/bookings/internal/helpers"
	"github.com/prakasht9/bookings/internal/models"
	"github.com/prakasht9/bookings/internal/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {

	gob.Register(models.Reservation{})
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t: ", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t: ", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	tc, err := render.CreateTemplateCache()
	if err != nil {

		log.Fatal("can not create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates((&app))
	helpers.NewHelpers(&app)
	return nil
}
