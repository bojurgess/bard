package main

import (
	"github.com/bojurgess/bard/internal/config"
	"github.com/bojurgess/bard/internal/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	c := config.AppConfig
	r := router.SetupRoutes()

	addr := c.Host + ":" + c.Port

	log.Println("Listening for requests at http://" + addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting server", err)
	}
}
