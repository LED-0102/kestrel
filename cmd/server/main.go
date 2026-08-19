package main

import (
	"log"

	"github.com/LED-0102/kestrel/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
