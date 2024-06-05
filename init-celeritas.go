package main

import (
	"log"
	"os"

	"github.com/jamesyang124/celeritas"
)

type application struct {
	App *celeritas.Celeritas
}

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init Celeritas type struct
	cel := &celeritas.Celeritas{}

	// runner to create Celeritas application, scaffold generator
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "celeritasApp"

	// ensure info logger is set as well
	cel.InfoLog.Println("debug mode is set by env:", cel.Debug)

	app := &application{
		App: cel,
	}

	return app
}
