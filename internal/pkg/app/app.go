package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/beardedandnotmuch/google-sheets-observer/internal/app/endpoint"
	"github.com/beardedandnotmuch/google-sheets-observer/internal/app/service"
	"github.com/joho/godotenv"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	http.HandleFunc("/api/sheets-observer", a.e.HandleClientRequest)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Listenning on port", os.Getenv("PORT"), ".")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running")

	return nil
}
