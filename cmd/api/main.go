package main

import (
	"github.com/bojurgess/bard/internal/config"
	"github.com/bojurgess/bard/internal/database"
	"github.com/bojurgess/bard/internal/router"
	"log"
	"net/http"
)

func main() {
	var err error

	config.Load()
	err = database.Initialize()
	c := config.AppConfig
	r := router.SetupRoutes()

	if err != nil {
		log.Fatal(err)
	}

	addr := c.Host + ":" + c.Port

	log.Println("Listening for requests at http://" + addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting server", err)
	}
}
