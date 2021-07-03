package main

import (
	"log"

	"github.com/quyenphamkhac/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
