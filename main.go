package main

import (
	"charlex-web-os/app"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("CharleX WebOS listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(app.Handler)))
}
