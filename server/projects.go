package server

import (
	"net/http"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}
