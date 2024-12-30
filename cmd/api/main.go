package main

import (
	"github.com/bojurgess/bard/internal/config"
	"log"
	"net/http"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	addr := c.Host + ":" + c.Port

	log.Println("Listening for requests at http://" + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Error starting server", err)
	}
}
