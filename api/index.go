package handler

import (
	"charlex-web-os/app"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.Handler(w, r)
}
