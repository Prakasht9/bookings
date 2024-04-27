package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prakasht9/bookings/internal/config"
	"github.com/prakasht9/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/general_suite", handlers.Repo.Generals)
	mux.Get("/luxury_suite", handlers.Repo.Majors)
	mux.Post("/search_availability", handlers.Repo.PostAvailability)
	mux.Post("/search_availability-json", handlers.Repo.AvailabilityJson)

	mux.Get("/search_availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make_reservation", handlers.Repo.Reservation)
	mux.Post("/make_reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation_summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
