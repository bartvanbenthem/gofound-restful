package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/bartvanbenthem/gofound-restful/internal/driver"
	"github.com/bartvanbenthem/gofound-restful/internal/handlers"
)

const portNumber = ":4000"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	log.Printf(fmt.Sprintf("Staring application on port %s\n", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// connect to database
	log.Printf("Connecting to database...\n")
	db, err := driver.ConnectSQL("postgres://postgres:password@localhost/go_software?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database: Fatal...")
	}

	log.Printf("Connected to database\n")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	return db, nil
}
