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
	cel.Debug = true

	app := &application{
		App: cel,
	}

	return app
}
