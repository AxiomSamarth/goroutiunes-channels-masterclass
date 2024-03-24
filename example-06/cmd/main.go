package main

import (
	"log"

	"github.com/AxiomSamarth/gcm/example-06/cmd/app"
)

func main() {
	service, err := app.NewService()
	if err != nil {
		panic(err)
	}

	log.Println("starting the service")
	service.Run()
}
