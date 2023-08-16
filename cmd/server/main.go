package main

import (
	"log"

	"github.com/alexjaw/distrib-services-go-book/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}